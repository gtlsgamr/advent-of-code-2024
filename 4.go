package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./inputs/4")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	str := string(content)
	lines := strings.Split(str, "\n")
	matrix := make([][]string, 0, len(lines))
	for _, line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
	}
	totalCount := 0
	totalMASCount := 0
	for r, line := range matrix {
		for c := range line {
			if matrix[r][c] == "X" {
				totalCount += calculateXmasFromX(matrix, r, c)
			}
			if matrix[r][c] == "A" {
				totalMASCount += calculateMASFromA(matrix, r, c)
			}
		}
	}

	fmt.Println(totalCount)
	fmt.Println(totalMASCount)

}

func calculateMASFromA(rows [][]string, r, c int) int {
	if r == 0 || c == 0 || r == len(rows)-1 || c == len(rows[0])-1 {
		return 0
	}

	count := 0

	// Case 1: M A S (top-left to bottom-right) and M A S (top-right to bottom-left)
	if (rows[r+1][c+1] == "S" && rows[r-1][c-1] == "M") &&
		(rows[r+1][c-1] == "S" && rows[r-1][c+1] == "M") {
		count++
	}

	// Case 2: S A M (top-left to bottom-right) and S A M (top-right to bottom-left)
	if (rows[r+1][c+1] == "M" && rows[r-1][c-1] == "S") &&
		(rows[r+1][c-1] == "M" && rows[r-1][c+1] == "S") {
		count++
	}

	// Case 3: M A S (top-left to bottom-right) and S A M (top-right to bottom-left)
	if (rows[r+1][c+1] == "S" && rows[r-1][c-1] == "M") &&
		(rows[r+1][c-1] == "M" && rows[r-1][c+1] == "S") {
		count++
	}

	// Case 4: S A M (top-left to bottom-right) and M A S (top-right to bottom-left)
	if (rows[r+1][c+1] == "M" && rows[r-1][c-1] == "S") &&
		(rows[r+1][c-1] == "S" && rows[r-1][c+1] == "M") {
		count++
	}

	return count
}

func calculateXmasFromX(rows [][]string, r, c int) int {
	count := calculateVertical(rows, r, c) + calculateHorizontal(rows, r, c) + calculateDiagonal(rows, r, c)
	return count
}

func calculateHorizontal(rows [][]string, r, c int) int {
	count := 0
	width := len(rows[0])

	// Check right (LTR)
	if c+3 < width &&
		rows[r][c+1] == "M" &&
		rows[r][c+2] == "A" &&
		rows[r][c+3] == "S" {
		count++
	}

	// Check left (RTL)
	if c-3 >= 0 &&
		rows[r][c-1] == "M" &&
		rows[r][c-2] == "A" &&
		rows[r][c-3] == "S" {
		count++
	}

	return count
}

func calculateVertical(rows [][]string, r, c int) int {
	count := 0
	height := len(rows)

	// Check down
	if r+3 < height &&
		rows[r+1][c] == "M" &&
		rows[r+2][c] == "A" &&
		rows[r+3][c] == "S" {
		count++
	}

	// Check up
	if r-3 >= 0 &&
		rows[r-1][c] == "M" &&
		rows[r-2][c] == "A" &&
		rows[r-3][c] == "S" {
		count++
	}

	return count
}

func calculateDiagonal(rows [][]string, r, c int) int {
	count := 0
	height := len(rows)
	width := len(rows[0])

	// Check diagonal down-right
	if r+3 < height && c+3 < width &&
		rows[r+1][c+1] == "M" &&
		rows[r+2][c+2] == "A" &&
		rows[r+3][c+3] == "S" {
		count++
	}

	// Check diagonal down-left
	if r+3 < height && c-3 >= 0 &&
		rows[r+1][c-1] == "M" &&
		rows[r+2][c-2] == "A" &&
		rows[r+3][c-3] == "S" {
		count++
	}

	// Check diagonal up-right
	if r-3 >= 0 && c+3 < width &&
		rows[r-1][c+1] == "M" &&
		rows[r-2][c+2] == "A" &&
		rows[r-3][c+3] == "S" {
		count++
	}

	// Check diagonal up-left
	if r-3 >= 0 && c-3 >= 0 &&
		rows[r-1][c-1] == "M" &&
		rows[r-2][c-2] == "A" &&
		rows[r-3][c-3] == "S" {
		count++
	}

	return count
}
