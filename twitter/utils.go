package twitter

// MinusIDList Return a list of id composed of a list id minus b list id
func MinusIDList(a []int64, b []int64) []int64 {
	bMap := make(map[int64]bool)
	for _, id := range b {
		bMap[id] = true
	}

	result := []int64{}
	for _, id := range a {
		if !bMap[id] {
			result = append(result, id)
		}
	}
	return result
}
