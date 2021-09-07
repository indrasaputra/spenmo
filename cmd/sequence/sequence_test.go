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
			[]int{2},
			[]int{1, 2},
			[]int{2, 3, 4},
			[]int{3, 4, 5, 6, 7},
			[]int{4, 5, 6, 7, 8, 9, 10},
			[]int{7, 8, 9},
			[]int{9, 10},
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
			[]int{1, 2, 4},
			[]int{2, 3, 4, 6},
			[]int{3, 4, 5, 6, 7, 9},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			[]int{7, 8, 9, 10, 11},
			[]int{9, 11},
			[]int{},
		}

		for _, seq := range seqs {
			exists := sequence.SequenceExists(numbers, seq)
			if exists {
				t.Error(fmt.Sprintf("seq %v should not exists!", seq))
			}
		}
	})
}
