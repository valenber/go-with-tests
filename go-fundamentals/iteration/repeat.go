package iteration

const defaultRepeatCount = 5

func Repeat(char string, repeat int) string {
	repeatCount := repeat
	var result string

	if repeatCount == 0 {
		repeatCount = defaultRepeatCount
	}

	for i := 0; i < repeatCount; i++ {
		result += char
	}

	return result
}
