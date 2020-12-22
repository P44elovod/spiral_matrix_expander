package util

import (
	"testing"

	"gopkg.in/stretchr/testify.v1/assert"
)

func Test_isMatrixSquare(t *testing.T) {
	testCases := []struct {
		name     string
		matrix   [][]int
		expected bool
	}{
		{
			name: "wrong",
			matrix: [][]int{
				{1, 2, 3, 1},
				{4, 5, 6, 4},
				{7, 8, 9, 7},
				{7, 8, 9},
			},
			expected: false,
		},
		{
			name: "right",
			matrix: [][]int{
				{1, 2, 3, 1},
				{4, 5, 6, 4},
				{7, 8, 9, 7},
				{7, 8, 9, 8},
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := isMatrixSquare(tc.matrix)
			assert.Equal(t, tc.expected, res)

		})

	}

}

func Test_expandMatrixClockwise(t *testing.T) {
	testCases := []struct {
		name     string
		matrix   [][]int
		expected []int
	}{
		{
			name: "first",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "second",
			matrix: [][]int{
				{1, 2, 3, 1},
				{4, 5, 6, 4},
				{7, 8, 9, 7},
				{7, 8, 9, 7},
			},
			expected: []int{1, 2, 3, 1, 4, 7, 7, 9, 8, 7, 7, 4, 5, 6, 9, 8},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := expandMatrixClockwise(tc.matrix)
			assert.Equal(t, tc.expected, res)
		})
	}
}
