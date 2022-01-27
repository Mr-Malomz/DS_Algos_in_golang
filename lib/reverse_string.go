package lib

func ReverseString(text string) string {
	var output string

	for i := len(text) - 1; i >= 0; i-- {
		output += output + string(text[i])
	}

	return output
}

func ReverseString2(text string) string {
	var output string
	for _, v := range text {
		output = string(v) + output
	}

	return output
}
