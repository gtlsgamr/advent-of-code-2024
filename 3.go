package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./inputs/3")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	defer file.Close()
	mulExp := regexp.MustCompile(`(do|don't)\(\)|mul\((\d{1,3}),(\d{1,3})\)`)
	//mulExp := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	scanner := bufio.NewScanner(file)
	resultSum1 := 0
	resultSum2 := 0
	var lastWasDo bool
	for scanner.Scan() {
		line := scanner.Text()
		matches := mulExp.FindAllString(line, -1)
		for _, m := range matches {
			switch m {
			case "do()":
				lastWasDo = true
			case "don't()":
				lastWasDo = false
			default:
				resultSum2 += mulFromStr(m, lastWasDo)
				resultSum1 += mulFromStr(m, true)
			}
		}
	}
	fmt.Println(resultSum1)
	fmt.Println(resultSum2)
}

func mulFromStr(str string, lastWasDo bool) int {
	valExp := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := valExp.FindAllStringSubmatch(str, -1)

	var sum int
	for _, match := range matches {
		val1, _ := strconv.Atoi(match[1])
		val2, _ := strconv.Atoi(match[2])

		if lastWasDo {
			sum += val1 * val2
		}
		// if lastWasDo is false (meaning last was don't), we skip the multiplication
	}

	return sum
}
