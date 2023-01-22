package utils

import (
	"io/ioutil"
	"os"
)

func LoadFile(filename string) ([]byte, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	return byteValue, nil
}
