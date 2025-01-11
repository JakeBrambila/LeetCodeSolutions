package main

import "fmt"

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	floodFill(matrix, 1, 1, 2)
	fmt.Println("After Flood fill: ")
	fmt.Println(matrix)
}

// this uses depth first search(DFS) recursive algorithm
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}

	fill(image, sr, sc, image[sr][sc], color)
	return image
}

// helper function to change one pixel
// makes recursive calls
func fill(image [][]int, sr int, sc int, color int, newColor int) {

	// base case
	// for index out of bounds checks ALWAYS check the indices BEFORE
	if sr < 0 || sc < 0 || sr >= len(image) || sc >= len(image[0]) || image[sr][sc] != color {
		return
	}

	image[sr][sc] = newColor
	fill(image, sr-1, sc, color, newColor)
	fill(image, sr+1, sc, color, newColor)
	fill(image, sr, sc-1, color, newColor)
	fill(image, sr, sc+1, color, newColor)
}
