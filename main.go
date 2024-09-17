package main

import (
	"bufio"
	"fmt"
	"os"
	"pairs/pairs"
	"strconv"
	"strings"
)

// run is the main entry-point to the program which gets input from user, sanitizes input, calls
// the main algorithm, or returns appropriate error.
func run() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a list of integers (space-separated):")
	input, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("reading list from stdin: %v", err)
	}

	input = strings.TrimSpace(input)
	numsStr := strings.Split(input, " ")

	// Convert the input string slice to an integer slice.
	var numbers []int
	for _, str := range numsStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			return fmt.Errorf("invalid input. Enter only integer values")
		}
		numbers = append(numbers, num)
	}

	fmt.Println("Enter the target integer:")
	targetInput, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("reading target from stdin: %v", err)
	}

	targetInput = strings.TrimSpace(targetInput)

	target, err := strconv.Atoi(targetInput)
	if err != nil {
		return fmt.Errorf("invalid input. Enter only integer values")
	}

	result := pairs.Find(numbers, target)
	pairs.Print(result)

	return nil
}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}
