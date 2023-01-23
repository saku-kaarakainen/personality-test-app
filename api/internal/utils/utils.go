package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Loader interface {
	LoadFile(filename string) ([]byte, error)
}
type FileLoader struct{}

func (f FileLoader) LoadFile(filename string) ([]byte, error) {
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

// Rips out the value from the input based by the template
// Works with one parameter "%s".
// Example:
//
//	template = "<p>%s</p>"
//	input = "<p>Hello</p>"
//	returns: "Hello"
func Unformat(template string, input string) (string, error) {
	var result strings.Builder
	// 37 = %
	// 115 = s
	sep_bytes := [2]byte{37, 115}

	// The logs are left in the comments for purpose, in case you need to debug this.
	for i := 0; i < len(input); i++ {
		// byte comparison
		if input[i] == template[i] {
			continue
		}

		// compare if we got %s
		if template[i] != sep_bytes[0] || template[i+1] != sep_bytes[1] {
			return "", fmt.Errorf("unable to parse from input '%s' with template '%s'", input, template)
		}

		template_index := i + 2

		for ; i < len(input); i++ {
			// check if we are again matching with the template
			if input[i] == template[template_index] {
				break
			}

			result.WriteByte(input[i])
		}

		// break already, because i was altered so it might be already pointing at end
		// So the next is ignored
		// TODO: take the rest into account.
		break
	}

	return result.String(), nil
}
