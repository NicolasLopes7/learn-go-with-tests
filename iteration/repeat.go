package iteration

func Repeat(char string, count int) (repeatedChar string) {
	for i := 0; i < count; i++ {
		repeatedChar += char
	}
	return
}
