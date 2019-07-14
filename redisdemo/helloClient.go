package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"log"
)

func connect_redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "password123",
		DB:       0, // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(pong, err)
		panic(err)
	}

	return client
}

func main() {
	client := connect_redis()

	err := client.Set("name", "wanglei", 0).Err() //忽略错误
	if err != nil {
		panic(err)
	}

	val, err := client.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

}
