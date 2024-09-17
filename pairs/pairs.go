// Package pairs contains implementation of an algorithm for solving the finding-all-two-sum-pairs problem.
package pairs

import "fmt"

// Find receives a list of integers and a target integer. It returns all pairs whose sum is
// equal to the given target.
func Find(numbers []int, target int) [][2]int {
	// This map holds the frequencies of occurrences of numbers in the list.
	// Frequencies are stored to handle duplicate input values as described in README.
	frequencies := make(map[int]int)

	// Populate the frequency map.
	for _, num := range numbers {
		frequencies[num]++
	}

	var result [][2]int

	for currentNumber, currentNumberFrequency := range frequencies {
		complement := target - currentNumber
		complementFrequency := frequencies[complement]

		if complementFrequency > 0 {
			var pairCount int
			if currentNumber == complement {
				// Special case when num and complement are the same in form of (num, num). So we
				// use combination formula for selecting 2 out.
				pairCount = currentNumberFrequency * (currentNumberFrequency - 1) / 2
			} else {
				pairCount = currentNumberFrequency * complementFrequency
			}

			for i := 0; i < pairCount; i++ {
				result = append(result, [2]int{currentNumber, complement})
			}

			delete(frequencies, currentNumber)
			delete(frequencies, complement)
		}
	}

	return result
}

// Print is a helper function that receives all pairs and prints a human-readable text to stdout.
func Print(pairs [][2]int) {
	if len(pairs) == 0 {
		fmt.Println("no pairs found")
		return
	}

	for _, pair := range pairs {
		fmt.Printf("(%d, %d)\n", pair[0], pair[1])
	}
}
