// Another apporach to try and examine: go from front and back of the string. One from the front, one from the back etc. This could enchance algorithms efficiency.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Prompt the user for input
	fmt.Print("Enter a string: ")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Call the function with the user's input
	sum := calculateSumOfNumbersInString(input)
	fmt.Println("Sum of numbers in the string:", sum)
}

func calculateSumOfNumbersInString(input string) int {
	stringArr := splitString(input, ' ')
	sum := sumCreatedNumbersOfStrings(stringArr)
	return sum
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
