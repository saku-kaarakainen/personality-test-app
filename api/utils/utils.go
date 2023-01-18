package utils

import (
	"fmt"
	"strings"
)

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

	for i := 0; i < len(input); i++ {
		// log.Printf("Comparing '%s' with '%s' at index '%d'.\n", string(input[i]), string(template[i]), i)
		// byte comparison
		if input[i] == template[i] {
			continue
		}

		// compare if we got %s
		if template[i] != sep_bytes[0] || template[i+1] != sep_bytes[1] {
			return "", fmt.Errorf("unable to parse from input '%s' with template '%s'", input, template)
		}

		// log.Println("Going to inner loop")
		template_index := i + 2

		for ; i < len(input); i++ {
			// log.Printf("Comparing '%s' with '%s' at index '%d'.\n", string(input[i]), string(template[template_index]), i)
			// check if we are again matching with the template
			if input[i] == template[template_index] {
				break
			}

			result.WriteByte(input[i])
			// log.Printf("constructed so far: %s\n", result.String())
		}

		// break already, because i was altered so it might be already pointing at end
		// So the next is ignored
		// TODO: take the rest into account.
		break
	}

	return result.String(), nil
}
