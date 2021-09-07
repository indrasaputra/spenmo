package sequence

// SequenceExists checks whether seq is found in numbers.
func SequenceExists(numbers, seq []int) bool {
	n := len(numbers)
	s := len(seq)

	if s > n {
		return false
	}
	if s == 0 {
		return false
	}

	for i := 0; i < n; i++ {
		if numbers[i] == seq[0] && isContainSeq(numbers[i:], seq) {
			return true
		}
	}
	return false
}

func isContainSeq(numbers, seq []int) bool {
	n := len(numbers)
	s := len(seq)
	if s > n {
		return false
	}

	for i := 0; i < s; i++ {
		if numbers[i] != seq[i] {
			return false
		}
	}
	return true
}
