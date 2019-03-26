package main

import "fmt"

func main() {
	err = redisClient.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisClient.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

}
