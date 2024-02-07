package iteration

func Repeat(input string, repeat int) (repated string) {
	for i := 0; i < repeat; i++ {
		repated += input
	}
	return
}
