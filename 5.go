package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./inputs/5")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	strContent := string(content)
	splitContent := strings.Split(strContent, "\n\n")
	mappings := strings.Split(splitContent[0], "\n")
	inputs := strings.Split(splitContent[1], "\n")

	var beforeMap = make(map[string]map[string]bool)
	var afterMap = make(map[string]map[string]bool)

	for _, mapping := range mappings {
		vals := strings.Split(mapping, "|")

		// For afterMap
		if _, ok := afterMap[vals[0]]; !ok {
			afterMap[vals[0]] = make(map[string]bool)
		}
		afterMap[vals[0]][vals[1]] = true

		// For beforeMap
		if _, ok := beforeMap[vals[1]]; !ok {
			beforeMap[vals[1]] = make(map[string]bool)
		}
		beforeMap[vals[1]][vals[0]] = true
	}

	var sum int
	var sum2 int
	var status bool

	for _, input := range inputs {
		inputArray := strings.Split(input, ",")
		status = calculateOrders(beforeMap, afterMap, inputArray)
		if status {
			val, err := strconv.Atoi(inputArray[len(inputArray)/2])
			if err != nil {
				log.Fatal(err)
			}
			sum += val
		} else {
			orderedArray := orderArray(beforeMap, afterMap, inputArray)
			middleIndex := len(orderedArray) / 2
			val, err := strconv.Atoi(orderedArray[middleIndex])
			if err != nil {
				log.Fatal(err)
			}
			sum2 += val
		}
	}

	fmt.Println(sum)
	fmt.Println(sum2)

}

// PART2
func orderArray(beforeMap, afterMap map[string]map[string]bool, inputArray []string) []string {
	n := len(inputArray)
	result := make([]string, n)
	copy(result, inputArray)

	// Bubble sort with custom comparison
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// If current element should come after next element
			if !checkNumber(afterMap, []string{result[j+1]}, result[j]) {
				// Swap elements
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}

	return result
}

// PART1
func calculateOrders(beforeMap map[string]map[string]bool, afterMap map[string]map[string]bool, input []string) bool {
	begIndex := 0
	endIndex := len(input) - 1
	for begIndex < endIndex {
		result1 := checkNumber(afterMap, input[begIndex+1:], input[begIndex])
		if !result1 {
			return false
		}
		result2 := checkNumber(beforeMap, input[:endIndex-1], input[endIndex])
		if !result2 {
			return false
		}
		begIndex = begIndex + 1
		endIndex = endIndex - 1
	}
	return true
}

func checkNumber(store map[string]map[string]bool, input []string, num string) bool {
	for _, val := range input {
		if _, ok := store[num][val]; !ok {
			return false
		}
	}
	return true
}
