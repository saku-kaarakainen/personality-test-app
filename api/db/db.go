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
	log.Println("Init mod db")
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

// questions
type Answer struct {
	Score [2]int `json:"id"`
	Label string `json:"question_label"`
}

type Question struct {
	Id      string   `json:"id"`
	Label   string   `json:"question_label"`
	Answers []Answer `json:"answers"`
}

// results
type Result struct {
	Id                    string   `json:"id"`
	Label                 string   `json:"question_label"`
	DescriptionParagraphs []string `json:"description_paragraphs"`
}

func setFromFile(filename string, key string, objType interface{}) error {
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

	res, err := rh.JSONSet(key, ".", objType)
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

func populate() {

	var questions []Question
	setFromFile("db/data/questions.json", "questions", questions)

	var results []Result
	setFromFile("db/data/results.json", "results", results)
}

func GetGuestions() ([]Question, error) {
	jsonBlob, err := rh.JSONGet("questions", ".")
	if err != nil {
		return nil, err
	}

	var data []Question
	json.Unmarshal(jsonBlob.([]byte), &data)

	return data, nil
}
