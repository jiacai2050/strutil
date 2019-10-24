package strutil

func Reverse(s string) string {
	runes := []rune(s)

	for i := 0; i < len(runes)/2; i++ {
		j := len(runes) - i - 1
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
