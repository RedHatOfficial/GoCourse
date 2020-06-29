package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

// adresa určující službu Redisu, která se má použít
const redisAddress = "localhost:6379"

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

	// přečtení hodnoty z databáze Redisu
	address, err := client.Get(context, "Seriál o jazyku Go").Result()

	// vyhodnocení předchozí operace
	switch {
	case err == redis.Nil:
		fmt.Println("no value found")
	case err != nil:
		panic(err)
	default:
		fmt.Println("Adresa:", address)
	}
}
