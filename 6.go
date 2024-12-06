package main

import (
	"fmt"
	"github.com/k0kubun/pp/v3"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./inputs/6")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	str := string(content)
	var rows [][]string
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		rows = append(rows, (strings.Split(line, "")))
	}
	var initialPositionRow, initialPositionCol int
	for i, line := range lines {
		if pos := strings.Index(line, "^"); pos != -1 {
			initialPositionRow = i
			initialPositionCol = pos
			break
		}
	}
	initialDirection := 0 // Directions will be 0,1,2,3 = up right down left

	pp.Print(initialPositionCol, initialPositionRow)

	status, distinct := moveUntilObstacle(rows, initialPositionRow, initialPositionCol, initialDirection, map[int]map[int]bool{
		initialPositionCol: {
			initialPositionRow: true,
		},
	})
	count := 0
	for _, colMap := range distinct {
		count += len(colMap)
	}
	fmt.Println(status, count)
}

func moveUntilObstacle(matrix [][]string, row, col int, direction int, distinct map[int]map[int]bool) (bool, map[int]map[int]bool) {
	if col+1 > len(matrix[0])-1 || row+1 > len(matrix)-1 || row-1 < 0 || col-1 < 0 {
		return true, distinct
	}

	var nextBlock string
	switch direction {
	case 0:
		nextBlock = matrix[row-1][col]
		if nextBlock != "#" {
			row = row - 1
		} else {
			direction = 1
		}
	case 1:
		nextBlock = matrix[row][col+1]
		if nextBlock != "#" {
			col = col + 1
		} else {
			direction = 2
		}
	case 2:
		nextBlock = matrix[row+1][col]
		if nextBlock != "#" {
			row = row + 1
		} else {
			direction = 3
		}
	case 3:
		nextBlock = matrix[row][col-1]
		if nextBlock != "#" {
			col = col - 1
		} else {
			direction = 0
		}
	}

	if _, exists := distinct[row]; !exists {
		distinct[row] = make(map[int]bool)
	}
	distinct[row][col] = true

	return moveUntilObstacle(matrix, row, col, direction, distinct)
}
