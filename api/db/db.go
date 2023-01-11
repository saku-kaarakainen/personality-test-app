package db

import (
	"context"
	"fmt"
	"log"

	goredis "github.com/go-redis/redis/v8" // uses redis7
	"github.com/gomodule/redigo/redis"
	"github.com/nitishm/go-rejson/v4"
)

type Test struct {
	Foo     string `json:"foo,omitempty"`
	SomeKey string `json:"some_key,omitempty"`
}

var (
	ctx = context.Background()
	cli = goredis.NewClient(&goredis.Options{
		// TODO: Add username and password, also .env
		Addr:     "personality-test-db:6379",
		Password: "", // no password set
		DB:       0,  // use default db
		//Network:     "db-network",
		//DialTimeout: 60,
	})
	rh = rejson.NewReJSONHandler()
)

func init() {
	rh.SetGoRedisClient(cli)

	ping()
	populate()
}

func ping() {
	pong, err := cli.Ping(ctx).Result()
	fmt.Println("Redis Ping:")
	fmt.Println(pong, err)

	if err != nil {
		panic(err)
	}
}

func populate() {
	fmt.Println("starting to populate")
	test := Test{
		Foo:     "hardcoded1",
		SomeKey: "hardcoded2",
	}

	res, err := rh.JSONSet("test", ".", test)
	if err != nil {
		log.Fatalf("Failed to JSONSet")
		return
	}

	fmt.Println("JSON with key 'test' set!")

	if res.(string) == "OK" {
		fmt.Printf("Success: %s\n", res)
	} else {
		fmt.Println("Failed to Set: ")
	}

	testJSON, err := redis.Bytes(rh.JSONGet("test", "."))
	if err != nil {
		fmt.Println("Failed to JSONGet")
		panic("Failed to JSONGet")
	}

	fmt.Println(("found json: " + string(testJSON)))
}

func LoadModule() {
	// makes sure init is called.
	// init is lazy-load so it will be called
	// when the module is referred first time.
}
