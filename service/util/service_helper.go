package util

func NewSkipNumber(currentNum, contentCount int) int {
	return currentNum * contentCount
}

func ConvertSliceToExistMap(s []string) map[string]int {
	m := make(map[string]int)
	for _, v := range s {
		m[v] = 1
	}
	return m
}
