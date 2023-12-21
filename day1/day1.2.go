// Another apporach to try and examine: go from front and back of the string. One from the front, one from the back etc. This could enchance algorithms efficiency.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	processedInput := readFile("input.txt")

	stringArr := splitString(processedInput, ' ')
	sum := sumCreatedNumbersOfStrings(stringArr)
	fmt.Print(sum)
}

func readFile(fileName string) string {
	// Open the input file
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return "err"
	}
	defer file.Close()

	// Create a scanner to read from the file
	scanner := bufio.NewScanner(file)

	// Create an empty slice to store the lines
	var lines []string

	// Read lines from the file
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Close the scanner
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return "err"
	}

	// Join the lines with spaces and print the result
	result := strings.Join(lines, " ")
	fmt.Println(result)

	return result
}

func sumCreatedNumbersOfStrings(input []string) int {

	sum := 0

	for _, value := range input {

		str := createListOfNumbersInString(value)

		num := createNumberFromFirstAndLastNumber(str)

		sum += num
	}

	return sum
}

func createListOfNumbersInString(input string) []int {
	slice := make([]int, 0)

	for _, char := range input {
		if char >= '0' && char <= '9' {
			digit, err := strconv.Atoi(string(char))

			if err == nil {
				slice = append(slice, digit)
			}
		}
	}
	return slice
}

func createNumberFromFirstAndLastNumber(input []int) int {
	// Convert the first number to string
	first := strconv.Itoa(input[0])

	if len(input) < 2 {
		result, _ := strconv.Atoi(first + first)
		return result
	}

	// Convert the last number to string
	last := strconv.Itoa(input[len(input)-1])

	// Concatenate the strings and convert them back to an integer
	result, _ := strconv.Atoi(first + last)

	return result
}

func splitString(input string, delimiter rune) []string {
	// Use the strings.Split function to split the input string
	// into an array of elements based on the delimiter.
	elements := strings.Split(input, string(delimiter))

	return elements
}

// Idea

func createNumberFromFrontAndBack(input []int) int {
	if len(input) == 0 {
		return 0
	}

	// Initialize indices for the front and back of the slice
	front := 0
	back := len(input) - 1

	// Initialize variables to store the digits from the front and back
	frontDigit := input[front]
	backDigit := input[back]

	// Initialize variables to keep track of the multiplier for the front and back digits
	frontMultiplier := 1
	backMultiplier := 1

	// Initialize the result
	result := 0

	for front <= back {
		// Add the front digit to the result with the appropriate multiplier
		result += frontDigit * frontMultiplier

		// If the front and back indices are the same, break out of the loop
		if front == back {
			break
		}

		// Add the back digit to the result with the appropriate multiplier
		result += backDigit * backMultiplier

		// Update multipliers
		frontMultiplier *= 10
		backMultiplier *= 10

		// Move the indices
		front++
		back--
	}

	return result
}
