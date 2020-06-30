package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// adresa určující službu Redisu, která se má použít
const redisAddress = "localhost:6379"

// jméno hodnoty použité pro implementaci jednoduché fronty
const queueName = "fronta"

type redisSession struct {
	client  *redis.Client
	context context.Context
}

// printQueueLength vypíše aktuální délku fronty, samotná délka je přitom
// získána jiným způsobem (vložením prvku, použitím LLen atd.)
func printQueueLength(length int64) {
	fmt.Printf("Queue length after enqueuing is %d\n", length)
}

// mustEnqueue zajistí vložení prvku do fronty popř. pád aplikace v případě,
// kdy vložení není možné provést (Redis je odpojen atd.)
func mustEnqueueInteger(session redisSession, key string, value int) {
	fmt.Printf("Enqueuing %d into queue named '%s'\n", value, key)
	// přidání prvku do seznamu
	length, err := session.client.LPush(
		session.context, key, value).Result()
	if err != nil {
		panic(err)
	}
	printQueueLength(length)
}

func producer(session redisSession, key string, from int, to int) {
	// postupné vložení prvků do fronty
	for i := from; i < to; i++ {
		mustEnqueueInteger(session, queueName, i)
		time.Sleep(1 * time.Second)
	}
}

func consumer(session redisSession, key string, timeout time.Duration) {
	// přečtení všech hodnot z fronty
	for {
		// pokus o přečtení hodnoty z fronty
		keyValue, err := session.client.BRPop(
			session.context, timeout, queueName).Result()

		// vyhodnocení předchozí operace
		switch {
		case err == redis.Nil:
			fmt.Println("no value found")
			return
		case err != nil:
			panic(err)
		default:
			key := keyValue[0]
			value := keyValue[1]
			fmt.Printf(
				"Value dequed from queue named '%s': '%s'\n",
				key, value)
		}

		length := session.client.LLen(session.context, queueName).Val()
		printQueueLength(length)
		fmt.Println()
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

	session := redisSession{
		client:  client,
		context: context,
	}

	// pokus o klasický handshake typu PING-PONG
	_, err := client.Ping(context).Result()
	if err != nil {
		panic(err)
	}

	// smazání seznamu, pokud existoval
	client.Del(context, queueName)

	// spustíme producenta zpráv
	go producer(session, queueName, 0, 10)

	// je možné pustit více producentů
	// go producer(session, queueName, 11, 20)

	timeout, err := time.ParseDuration("10s")
	if err != nil {
		panic(err)
	}

	// nyní již můžeme spustit konzumenta zpráv
	consumer(session, queueName, timeout)
}
