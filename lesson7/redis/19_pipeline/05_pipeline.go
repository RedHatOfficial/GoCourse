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
	client.Del(context, "counter1")
	client.Del(context, "counter2")
	client.Del(context, "accumulator")

	var counter1 *redis.IntCmd
	var counter2 *redis.IntCmd
	var accumulator *redis.FloatCmd

	for i := 0; i < 5; i++ {
		_, err = client.Pipelined(context, func(pipe redis.Pipeliner) error {
			counter1 = pipe.Incr(context, "counter1")
			counter2 = pipe.Decr(context, "counter2")
			accumulator = pipe.IncrByFloat(context, "accumulator", 3.14)
			return nil
		})
		if err != nil {
			panic(err)
		}

		fmt.Printf("1st counter: %d\n", counter1.Val())
		fmt.Printf("2nd counter: %d\n", counter2.Val())
		fmt.Printf("accumulator: %f\n\n", accumulator.Val())
	}
}
