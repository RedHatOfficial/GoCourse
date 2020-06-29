package main

import (
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

	// zápis hodnoty do databáze Redisu
	err := client.Set(context,
		"Seriál o jazyku Go",
		"https://www.root.cz/serialy/programovaci-jazyk-go/",
		0).Err()
	if err != nil {
		panic(err)
	}
}
