package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// adresa určující službu Redisu, která se má použít
const redisAddress = "localhost:6379"

func mustPush(client *redis.Client, context context.Context, key string, value string) {
	fmt.Println("Pushing", value, "into", key)
	// přidání prvku do seznamu
	length, err := client.LPush(context, key, value).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("List length", length)
}

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

	// pokus o klasický handshake typu PING-PONG
	_, err := client.Ping(context).Result()
	if err != nil {
		panic(err)
	}

	// smazání seznamu, pokud existoval
	client.Del(context, "seznam")

	mustPush(client, context, "seznam", "foo")
	mustPush(client, context, "seznam", "bar")
	mustPush(client, context, "seznam", "baz")

	fmt.Println()

	// přečtení všech hodnot ze seznamu
	for {
		// pokus o přečtení hodnoty ze seznamu
		value, err := client.LPop(context, "seznam").Result()

		// vyhodnocení předchozí operace
		switch {
		case err == redis.Nil:
			fmt.Println("no value found")
			return
		case err != nil:
			panic(err)
		default:
			fmt.Println("Value from list", value)
		}

		length := client.LLen(context, "seznam").Val()
		fmt.Println("List length", length)
	}
}
