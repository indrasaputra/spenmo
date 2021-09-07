package sequence_test

import (
	"fmt"
	"testing"

	"github.com/indrasaputra/spenmo/cmd/sequence"
)

func TestSequenceExists(t *testing.T) {
	t.Run("sequence exists", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		seqs := [][]int{
			{2},
			{1, 2},
			{2, 3, 4},
			{3, 4, 5, 6, 7},
			{4, 5, 6, 7, 8, 9, 10},
			{7, 8, 9},
			{9, 10},
		}

		for _, seq := range seqs {
			exists := sequence.SequenceExists(numbers, seq)
			if !exists {
				t.Error(fmt.Sprintf("seq %v should exists!", seq))
			}
		}
	})

	t.Run("sequence doesn't exists", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		seqs := [][]int{
			{1, 2, 4},
			{2, 3, 4, 6},
			{3, 4, 5, 6, 7, 9},
			{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			{7, 8, 9, 10, 11},
			{9, 11},
			{},
		}

		for _, seq := range seqs {
			exists := sequence.SequenceExists(numbers, seq)
			if exists {
				t.Error(fmt.Sprintf("seq %v should not exists!", seq))
			}
		}
	})
}
