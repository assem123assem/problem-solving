package problem8

// time complexity is O(n)
// space complexity is O(n)
func FirstDuplicateIndex(arr []int) int {
	duplicateMap := make(map[int]int)
	for idx, value := range arr {
		if duplicateMap[value] == 1 {
			return idx
		} else {
			duplicateMap[value] = 1
		}
	}
	return -1
}
