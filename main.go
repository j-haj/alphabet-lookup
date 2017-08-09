// this package converts numbers to letters, where 1 = 'a'
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// convertToAlphbet takes a sequence of numbers less than 27
// representing the letters of the alphabet, and converts them
// to their corresponding letters
func convertToAlphbet(input []string) string {
	output := make([]byte, len(input))
	for idx, c := range input {
		val, err := strconv.Atoi(c)
		if err != nil {
			fmt.Printf("error - %s\n", err)
			return ""
		}
		// Add 96 to convert to ASCII code
		output[idx] = byte(val) + 96
	}
	return string(output)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Number to letter conversion tool")

	for {
		// Get input
		fmt.Printf("Enter sequence of numbers less than 26 (:quit to quit): ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error - %v\n", err)
			return
		}
		text = strings.TrimSpace(text)
		if text == ":quit" {
			break
		}

		// Get result
		sequence := strings.Fields(text)
		result := convertToAlphbet(sequence)
		fmt.Printf("Result: %s\n", result)
	}
}
