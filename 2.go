package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./inputs/2")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	defer file.Close()

	safeCount := 0
	dSafeCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result := strings.Split(line, " ")
		intList := make([]int, len(result))
		for i := 0; i < len(result); i++ {
			intList[i], _ = strconv.Atoi(result[i])
		}
		if determineLevel(intList) {
			safeCount++
		}
		if determineLevelDampener(intList) {
			dSafeCount++
		}
	}
	fmt.Println(safeCount)
	fmt.Println(dSafeCount)
}

func determineLevel(report []int) bool {
	if len(report) < 3 {
		return true
	}

	statusIncreasing := report[0] < report[1]
	// Check first step
	firstStep := abs(report[1] - report[0])
	if firstStep < 1 || firstStep > 3 {
		return false
	}

	// Check remaining sequence
	for i := 1; i < len(report)-1; i++ {
		// Check if adjacent numbers are equal
		if report[i] == report[i+1] {
			return false
		}

		currentStep := abs(report[i+1] - report[i])
		currentStatus := report[i] < report[i+1]

		// Check if direction changed
		if currentStatus != statusIncreasing {
			return false
		}

		// Check step size
		if currentStep < 1 || currentStep > 3 {
			return false
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func determineLevelDampener(report []int) bool {
	// First check if it's safe without removing any number
	if isSequenceSafe(report) {
		return true
	}

	// Try removing each number one at a time
	for i := 0; i < len(report); i++ {
		// Create a new slice without the current number
		newReport := make([]int, 0)
		newReport = append(newReport, report[:i]...)
		newReport = append(newReport, report[i+1:]...)

		// Check if removing this number makes the sequence safe
		if isSequenceSafe(newReport) {
			return true
		}
	}

	return false
}

func isSequenceSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}

	statusIncreasing := report[0] < report[1]
	// Check first step
	firstStep := abs(report[1] - report[0])
	if firstStep < 1 || firstStep > 3 {
		return false
	}

	// Check remaining sequence
	for i := 1; i < len(report)-1; i++ {
		// Check if adjacent numbers are equal
		if report[i] == report[i+1] {
			return false
		}

		currentStep := abs(report[i+1] - report[i])
		currentStatus := report[i] < report[i+1]

		// Check if direction changed
		if currentStatus != statusIncreasing {
			return false
		}

		// Check step size
		if currentStep < 1 || currentStep > 3 {
			return false
		}
	}

	return true
}
