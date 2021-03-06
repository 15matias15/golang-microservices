package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSortWorstCase(t *testing.T) {
	els := []int{9, 8, 7, 6, 5}
	els = BubbleSort(els)

	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}

func TestBubbleSortBestCase(t *testing.T) {
	els := []int{5, 6, 7, 8, 9}
	els = BubbleSort(els)

	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}

func getElements(n int) []int {
	results := make([]int, n)
	i := 0

	for j := n - 1; j >= 0; j-- {
		results[i] = j
		i++
	}
	return results
}

func TestElements(t *testing.T) {
	els := getElements(5)
	assert.NotNil(t, els)

	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 4, els[0])
	assert.EqualValues(t, 3, els[1])
	assert.EqualValues(t, 2, els[2])
	assert.EqualValues(t, 1, els[3])
	assert.EqualValues(t, 0, els[4])
}

func BenchmarkBubbleSort100000(b *testing.B) {
	els := getElements(100000)

	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}
