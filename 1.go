package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./inputs/1")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	defer file.Close()

	var list1, list2 []int32

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result := strings.Split(line, "   ")
		list1int, _ := strconv.Atoi(result[0])
		list2int, _ := strconv.Atoi(result[1])
		if len(result) >= 2 {
			list1 = append(list1, int32(list1int))
			list2 = append(list2, int32(list2int))
		} else {
			log.Printf("Warning: Line '%s' doesn't have two tab-separated values", line)
		}
	}

	// Check for scanner errors after the loop
	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning file:", err)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	var sum int32

	for i := 0; i < len(list1); i++ {
		var difference int32
		if list1[i] >= list2[i] {
			difference = list1[i] - list2[i]
		} else {
			difference = list2[i] - list1[i]
		}
		sum += difference
	}
	// Part One Answer
	fmt.Println(sum)

	rightListCounts := make(map[int32]int32)
	for _, num := range list2 {
		rightListCounts[num]++
	}

	var similarityScore int32
	for _, num := range list1 {
		similarityScore += num * rightListCounts[num]
	}

	fmt.Println(similarityScore)

}
