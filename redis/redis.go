package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func createConnection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return client
}

func setValue(client *redis.Client) {
	err := client.Set("name", "Gaurav", 0).Err()

	if err != nil {
		fmt.Println(err.Error)
	}
}

func getValue(client *redis.Client) {
	val, err := client.Get("name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}

func main() {
	client := createConnection()
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	setValue(client)
	getValue(client)
}
