package test

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Prompt the user for the file path
	fmt.Print("Enter the file path: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filePath := scanner.Text()

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialize the sum
	sum := 0

	// Regular expression to match digits or numbers in string form
	re := regexp.MustCompile(`(\d+)|(\w+)\s*=\s*(\d+)`)

	// Read the file line by line
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)

		// Iterate through matches and extract numbers
		for _, match := range matches {
			if match[1] != "" {
				// Found a digit, convert it to an integer
				digit, _ := strconv.Atoi(match[1])
				sum += digit
			} else if match[2] != "" && match[3] != "" {
				// Found a number in string form (e.g., "7 = seven")
				numberStr := match[3]
				digit, err := strconv.Atoi(numberStr)
				if err != nil {
					fmt.Printf("Error converting %s to an integer: %v\n", numberStr, err)
				} else {
					sum += digit
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the sum of numbers
	fmt.Printf("Sum of numbers: %d\n", sum)
}
