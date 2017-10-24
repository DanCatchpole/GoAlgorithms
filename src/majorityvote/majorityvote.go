package majorityvote

// perform the Boyer-Moore Majority Vote algorithm
func MajorityVote(arr []string) string {
	candidate := ""
	count := 0
	// perform the initial algorithm
	for _, v := range arr {
		if count == 0 {
			candidate = v
		}
		if candidate == v {
			count++
		} else {
			count--
		}
	}
	// perform a check to see if they are the majority
	totalCount := 0
	for _, v := range arr {
		if v == candidate {
			totalCount++
		}
	}
	if totalCount > len(arr)/2 {
		return candidate
	} else {
		return ""
	}
}
