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

	// pokus o klasický handshake typu PING-PONG
	_, err := client.Ping(context).Result()
	if err != nil {
		panic(err)
	}

	// smazání hodnoty, pokud existovala
	client.Del(context, "accumulator")

	// inkrementace (neexistující) hodnoty
	newValue := client.IncrByFloat(context, "accumulator", 0.0).Val()
	fmt.Println("Accumulator value:", newValue)

	// přečtení hodnoty z databáze Redisu
	newValue = client.IncrByFloat(context, "accumulator", 0.0).Val()
	fmt.Println("Accumulator value:", newValue)

	// inkrementace (nyní již existující) hodnoty
	newValue = client.IncrByFloat(context, "accumulator", 3.14).Val()
	fmt.Println("Accumulator value:", newValue)

	// dekrementace (nyní již existující) hodnoty
	newValue = client.IncrByFloat(context, "accumulator", -10e12).Val()
	fmt.Println("Accumulator value:", newValue)
}
