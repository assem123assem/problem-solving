package problem5

func Unique(wordsArr []string) []string {
	var uniqueList []string
	appearanceMap := make(map[string]int)
	for _, word := range wordsArr {
		appearanceMap[word]++
	}
	for word, count := range appearanceMap {
		if count == 1 {
			uniqueList = append(uniqueList, word)
		}
	}
	return uniqueList

}
