package util

import (
	"errors"
	"sort"
)

//ExpandSquareMatrix - check if matrix square, if it is - expand matrix clockwise and return one-dimencional slice, if not - return error
func ExpandSquareMatrix(sourceSlice [][]int) ([]int, error) {
	if isMatrixSquare(sourceSlice) {

		return expandMatrixClockwise(sourceSlice), nil

	}
	return nil, errors.New("Matrix is not square!")
}

func isMatrixSquare(matrix [][]int) bool {
	for i := range matrix {
		if len(matrix[i]) != len(matrix) {
			return false
		}
	}
	return true
}

func expandMatrixClockwise(sourceSlice [][]int) []int {
	var resultSlice []int
	for {

		if len(sourceSlice) == 0 {
			break
		}

		resultSlice = append(resultSlice, sourceSlice[0]...)

		for i := 1; i < len(sourceSlice); i++ {
			resultSlice = append(resultSlice, sourceSlice[i][len(sourceSlice)-1])
			sourceSlice[i] = sourceSlice[i][:len(sourceSlice)-1]
			sort.Slice(sourceSlice[i], func(i, j int) bool {
				return true
			})
		}
		sourceSlice = sourceSlice[1:]

		sort.Slice(sourceSlice, func(i, j int) bool {
			return true
		})

	}

	return resultSlice
}
