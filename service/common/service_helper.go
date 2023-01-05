package common

func NewSkipNumber(currentNum, contentCount int) int {
	return currentNum * contentCount
}

// ConvertSliceToExistMap slice 의 요소들을 map 에 담는다. map 은 slice 의 요소가 존재하는지(contain 여부) 확인할 때 사용한다.
func ConvertSliceToExistMap(s []string) map[string]int {
	m := make(map[string]int)
	for _, v := range s {
		m[v] = 1
	}
	return m
}
