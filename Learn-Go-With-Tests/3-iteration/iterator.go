package iteration

func Repeat(text string, times int) string {
	var result string
	for i := 0; i < times; i++ {
		result += text
	}
	return result
}
