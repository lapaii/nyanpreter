package util

func SplitSlice(slice []byte, split byte) [][]byte {
	output := [][]byte{}

	thisSlice := []byte{}
	for _, val := range slice {
		if val == split {
			output = append(output, thisSlice)
			thisSlice = []byte{}
			continue
		}

		thisSlice = append(thisSlice, val)
	}

	return output
}
