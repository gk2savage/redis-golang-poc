package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx       = context.Background()
	rdbClient *redis.Client
)

func main() {
	fmt.Println("Redis POC")

	rdbClient := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.62:6379",
		PoolSize: 1000,
		Password: "testpassword",
		DB:       0,
	})

	pong, err := rdbClient.Ping(ctx).Result()
	fmt.Println(pong, err)

	setter := rdbClient.Set(ctx, "file.exe", "a533e66ea4e448652e8a070f4609a384", 0).Err()
	if setter != nil {
		panic(setter)
	}
	setter1 := rdbClient.Set(ctx, "file.exe", "e332e66ea4e448652e8a070f4609a384", 345600000000000).Err()
	if setter != nil {
		panic(setter1)
	}
	setter2 := rdbClient.Set(ctx, "file.pdf", "98a1e66ea4e448652e8a070f4609a384", 0).Err()
	if setter != nil {
		panic(setter2)
	}

	fmt.Println("\nGetting Keys:")
	val, err := rdbClient.Get(ctx, "file.exe").Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("file.exe", val)
	}

	val2, err := rdbClient.Get(ctx, "virus.exe").Result()
	if err == redis.Nil {
		fmt.Println("virus.exe: key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("virus.exe", val2)
	}

	fmt.Println("\nListing all keys in Redis:")
	iter := rdbClient.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println(iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}

	fmt.Println("\nChecking if file.exe exists:")
	iter2 := rdbClient.Scan(ctx, 0, "*", 0).Iterator()
	for iter2.Next(ctx) {
		if iter2.Val() == "file.exe" {
			fmt.Println(iter2.Val(), "exists.")
		}
	}
	if err := iter2.Err(); err != nil {
		panic(err)
	}

	setter3 := rdbClient.GetSet(ctx, "cmd.exe", "cccce66ea4e448652e8a070f4609a384").Err()
	if setter3 != nil {
		panic(setter3)
	}
	fmt.Println("\nGetSet: Done")

	fmt.Println("\nTimeToLive")
	_, err = rdbClient.Expire(ctx, "cmd.exe", 96*time.Hour).Result()
	ttl, _ := rdbClient.TTL(ctx, "cmd.exe").Result()
	fmt.Println(ttl)

	if iskeyinRedis("count") {
		fmt.Println("Skipping... Key is present")
	} else {
		fmt.Println("Key is initiated")
	}
}

// func iskeyinRedis(key string) bool {

// 	rdbClient := redis.NewClient(&redis.Options{
// 		Addr:     "192.168.1.62:6379",
// 		Password: "testpassword",
// 		DB:       0,
// 	})

// 	val, _ := rdbClient.Incr(context.TODO(), key).Result()
// 	if val > 1 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func iskeyinRedis(key string) bool {

	rdbClient := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.62:6379",
		PoolSize: 1000,
		Password: "testpassword",
		DB:       0,
	})

	_, err := rdbClient.Get(context.TODO(), key).Result()
	if err == redis.Nil {
		setter := rdbClient.Set(context.TODO(), key, 1, 345600000000000).Err()
		fmt.Println(setter)
		return false
	} else {
		return true
	}
}
