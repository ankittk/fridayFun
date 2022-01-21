package main

import (
	"fmt"
)

func main() {
	if len(matrix1) == 0 || len(matrix1[0]) == 0 {
		fmt.Errorf("no maze to play with %v", matrix1)
		return
	}
	smallestPath := store(matrix1)
	fmt.Printf("Smallest Path: %v\n", smallestPath)

	smallestPath2 := store(matrix2)
	fmt.Printf("Smallest Path: %v\n", smallestPath2)

	smallestPath3 := store(matrix3)
	fmt.Printf("Smallest Path: %v\n", smallestPath3)
}

func recursion(matrix *[][]int, x, y int) int {
	if !isAtEndX(&(*matrix)[y], x) && !isAtEndY(matrix, y) {
		return (*matrix)[y][x] + min(recursion(matrix, x+1, y), recursion(matrix, x, y+1))
	}
	if !isAtEndX(&(*matrix)[y], x) {
		return (*matrix)[y][x] + recursion(matrix, x+1, y)
	}
	if !isAtEndY(matrix, y) {
		return (*matrix)[y][x] + recursion(matrix, x, y+1)
	}
	return (*matrix)[y][x]
}

func store(oldMatrix [][]int) int {
	m := len(oldMatrix[0])
	n := len(oldMatrix)
	matrix := make([][]int, len(oldMatrix))
	for y, _ := range oldMatrix {
		matrix[y] = make([]int, len(oldMatrix[y]))
		for x, _ := range oldMatrix[0] {
			if x == 0 && y == 0 {
				matrix[y][x] = oldMatrix[y][x]
				continue // identity of 0,0
			}
			if x == 0 {
				matrix[y][x] = matrix[y-1][x] + oldMatrix[y][x]
				continue
			}
			if y == 0 {
				matrix[y][x] = matrix[y][x-1] + oldMatrix[y][x]
				continue
			}
			matrix[y][x] = min(matrix[y-1][x], matrix[y][x-1]) + oldMatrix[y][x]
		}
	}
	return matrix[n-1][m-1]
}

func isAtEndX(row *[]int, x int) bool {
	return x == len(*row)-1
}

func isAtEndY(col *[][]int, y int) bool {
	return y == len(*col)-1
}

func isAtBase(matrix [][]int, x, y int) bool {
	return isAtEndX(&matrix[0], x) && isAtEndY(&matrix, y)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
