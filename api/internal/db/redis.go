package db

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	goredis "github.com/go-redis/redis/v8" // uses redis7
	"github.com/nitishm/go-rejson/v4"
)

type RedisDb struct {
	ctx context.Context
	cli *goredis.Client
	rh  *rejson.Handler

	logger log.Logger
}

func NewRedisDb(
	ctx context.Context,

	// configs
	addr string,
	password string,
	db int,
) *RedisDb {
	var (
		cli = goredis.NewClient(&goredis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		})
		rh = rejson.NewReJSONHandler()
	)

	rh.SetGoRedisClient(cli)

	return &RedisDb{
		ctx: ctx,
		cli: cli,
		rh:  rh,
	}
}

// Panics if the server does not ping
// when pinging, should answer with pong.
func (db *RedisDb) MustPing() {
	pong, err := db.cli.Ping(db.ctx).Result()
	db.logger.Println("Redis Ping:")
	db.logger.Println(pong, err)

	if err != nil {
		panic(err)
	}
}

// TODO: Break this method down into 2 separate things:
//  1. open file and store it to a byte array
//  2. store byte array into redis store.
func (db *RedisDb) SetFromFile(filename string, key string, objType interface{}) error {
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open json file '%s'.\n", filename)
		return err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Failed to get bytes of json file '%s'.\n", filename)
		return err
	}

	json.Unmarshal(byteValue, &objType)

	res, err := db.rh.JSONSet(key, ".", objType)
	if err != nil {
		log.Fatalf("Failed to store to redis json bytes with key '%s'.\n", key)
		return err
	}

	if res.(string) != "OK" {
		db.logger.Fatalf(
			"Failed to store file '%s' was stored in Redis with key '%s'. The response '%s'.\n",
			filename,
			key,
			res,
		)
		return errors.New("Failed to store file was stored in Redis with key.\n")
	}

	db.logger.Printf(
		"File '%s' was stored in Redis with key '%s'. The response '%s'.\n",
		filename,
		key,
		res,
	)

	return nil
}

func (db *RedisDb) Get(key string, path string) ([]byte, error) {
	jsonBlob, err := db.rh.JSONGet(key, path)
	if err != nil {
		return nil, err
	}

	return jsonBlob.([]byte), nil
}
