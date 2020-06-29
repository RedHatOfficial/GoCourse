package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

// adresa určující službu Redisu, která se má použít
const redisAddress = "localhost:6379"

func main() {
	// konfigurační parametry nutné pro připojení k Redisu
	options, err := redis.ParseURL("redis://" + redisAddress)
	if err != nil {
		panic(err)
	}

	// vytvoření nového klienta s předáním konfiguračních parametrů
	client := redis.NewClient(options)

	// neměli bychom zapomenout na ukončení práce s klientem
	defer func() {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}()

	// pouze zobrazíme textovou podobu struktury Client
	fmt.Println("Redis client:", client)

	// získáme kontext a zobrazíme informace o něm
	context := client.Context()
	fmt.Println("Context:", context)

	// pokus o klasický handshake typu PING-PONG
	pong, err := client.Ping(context).Result()
	fmt.Println("Ping-pong result:", pong, err)
}
