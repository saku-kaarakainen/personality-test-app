package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	goredis "github.com/go-redis/redis/v8" // uses redis7
	"github.com/nitishm/go-rejson/v4"
)

/*
	TODO: Create integration test for the methods of this module.
	It is possible to do unit test to test each method,
	but the tests would be meaningless.

	It would be more meaningful to do integration test to test
	if the data is actually stored and fetched from the database.

	Once that test is implemented, Ping() - method does not make sense in the way it is implemented.
	Afterwards you need to reconsider the purpose of Ping(), whether to keep it or reason it's usage.
*/

type IDb interface {
	// Config / Admin
	Ping()
	Populate()

	// Questions
	GetGuestions() ([]Question, error)

	// Results
	GetPoint(key string, value string) ([2]int32, error)
	GetResult(score [2]int32) (Result, error)
}

type Db struct {
	ctx context.Context
	cli *goredis.Client
	rh  *rejson.Handler
}

func NewDb(
	ctx context.Context,
	cli *goredis.Client,
	rh *rejson.Handler,
) *Db {
	return &Db{
		ctx: ctx,
		cli: cli,
		rh:  rh,
	}
}

// DDD: Admin / Config

func (db *Db) Ping() {
	pong, err := db.cli.Ping(db.ctx).Result()
	fmt.Println("Redis Ping:")
	fmt.Println(pong, err)

	if err != nil {
		panic(err)
	}
}

func (db *Db) setFromFile(filename string, key string, objType interface{}) error {
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
		log.Fatalf(
			"Failed to store file '%s' was stored in Redis with key '%s'. The response '%s'.\n",
			filename,
			key,
			res,
		)
		return errors.New("Failed to store file was stored in Redis with key.\n")
	}

	log.Printf(
		"File '%s' was stored in Redis with key '%s'. The response '%s'.\n",
		filename,
		key,
		res,
	)

	return nil
}

func (db *Db) Populate() {
	var questions []Question
	db.setFromFile("db/data/questions.json", "questions", questions)

	var results []Result
	db.setFromFile("db/data/results.json", "results", results)
}

// DDD: Questions

func (db *Db) GetGuestions() ([]Question, error) {
	jsonBlob, err := db.rh.JSONGet("questions", ".")
	if err != nil {
		return nil, err
	}

	var data []Question
	json.Unmarshal(jsonBlob.([]byte), &data)

	return data, nil
}

// DDD: Result

func (db *Db) GetPoint(key string, value string) ([2]int32, error) {
	path := fmt.Sprintf("$.[?(@.id==\"%s\")].answers[?(@.id==\"%s\")].score", key, value)
	jsonBlob, err := db.rh.JSONGet("questions", path)
	if err != nil {
		log.Println("err:")
		log.Println(err)
		return [2]int32{0, 0}, err
	}

	var data [1][2]int32
	json.Unmarshal(jsonBlob.([]byte), &data)

	return data[0], nil
}

func (db *Db) GetResult(score [2]int32) (Result, error) {
	return Result{}, errors.New("not implemented GetResult")
}
