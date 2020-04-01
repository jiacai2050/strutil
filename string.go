package strutil

import "log"

func init() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
}

func Reverse(s string) string {
	log.Printf("reverse in v1...")
	runes := []rune(s)

	for i := 0; i < len(runes)/2; i++ {
		j := len(runes) - i - 1
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
