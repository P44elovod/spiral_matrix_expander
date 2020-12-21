package main

import (
	"fmt"
	"spiral_matrix_expander/util"
)

func main() {
	testSlice := [][]int{
		{1, 2, 3, 1},
		{4, 5, 6, 4},
		{7, 8, 9, 7},
		{7, 8, 9, 7},
	}

	result, err := util.ExpandSquareMatrix(testSlice)
	util.FailOnError(err, "Matrix expanding went wrong!")

	fmt.Println(result)

}
