package problem1

func IsPalindrome(word string) bool {
	var leftPtr = 0
	var rightPtr = len(word) - 1

	for leftPtr < rightPtr {
		if word[leftPtr] != word[rightPtr] {
			return false
		}
		leftPtr++
		rightPtr--
	}

	return true
}
