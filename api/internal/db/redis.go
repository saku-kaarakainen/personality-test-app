package db

import (
	"context"
	"errors"
	"log"

	goredis "github.com/go-redis/redis/v8" // uses redis7
	"github.com/nitishm/go-rejson/v4"

	"github.com/saku-kaarakainen/personality-test-app/api/internal/config"
)

type RedisDb struct {
	ctx context.Context
	cli *goredis.Client
	rh  *rejson.Handler
}

func NewRedisDb(
	ctx context.Context,
	cfg config.Config,
) *RedisDb {
	cli := goredis.NewClient(&goredis.Options{
		Addr:     cfg.Db.Addr,
		Password: cfg.Db.Pw,
		DB:       cfg.Db.SelectedDb,
	})

	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClient(cli)

	return &RedisDb{
		ctx: ctx,
		cli: cli,
		rh:  rh,
	}
}

// Panics if the server does not ping
// when pinging, should answer with pong.
func (db *RedisDb) Ping() (string, error) {
	pong, err := db.cli.Ping(db.ctx).Result()
	if err != nil {
		return "", nil
	}

	return pong, nil
}

func (db *RedisDb) Update(key string, object interface{}) error {
	res, err := db.rh.JSONSet(key, ".", object)
	if err != nil {
		log.Fatalf("Failed to store to redis json bytes with key '%s'.\n", key)
		return err
	}

	if res.(string) != "OK" {
		log.Fatalf("Failed to store in Redis with key '%s'. The response '%s'.\n", key, res)
		return errors.New("Failed to store file was stored in Redis with a key.\n")
	}

	log.Printf("Item was stored in Redis with key '%s'. The response '%s'.\n", key, res)

	return nil
}

func (db *RedisDb) Get(key string, path string) (interface{}, error) {
	jsonBlob, err := db.rh.JSONGet(key, path)
	if err != nil {
		return nil, err
	}

	return jsonBlob, nil
}
