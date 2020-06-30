package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// adresa určující službu Redisu, která se má použít
const redisAddress = "localhost:6379"

// jméno kanálu
const channelName = "c1"

// struktura obsahující "session"
type redisSession struct {
	client  *redis.Client
	context context.Context
}

// zdroj zpráv
func publisher(session redisSession, channel string, from int, to int) {
	for i := from; i < to; i++ {
		err := session.client.Publish(session.context, channel, i).Err()
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}

// příjemce zpráv
func subscriber(session redisSession, channel string) {
	pubsub := session.client.Subscribe(session.context, channel)
	ch := pubsub.Channel()

	for message := range ch {
		fmt.Printf("Channel: %s  Message: '%s'\n",
			message.Channel,
			message.Payload)
	}
}

// vstupní bod do demonstračního příkladu
func main() {
	// vytvoření nového klienta s předáním konfiguračních parametrů
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// neměli bychom zapomenout na ukončení práce s klientem
	defer func() {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}()

	// získáme kontext
	context := client.Context()

	// vytvoříme session
	session := redisSession{
		client:  client,
		context: context,
	}

	// pokus o klasický handshake typu PING-PONG
	_, err := client.Ping(context).Result()
	if err != nil {
		panic(err)
	}

	// smazání kanálu, pokud existoval
	client.Del(context, channelName)

	// nyní můžeme spustit několik příjemců zpráv
	go subscriber(session, channelName)
	go subscriber(session, channelName)
	go subscriber(session, channelName)

	// spustíme zdroj zpráv
	publisher(session, channelName, 0, 10)
}
