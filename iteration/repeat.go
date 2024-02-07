package iteration

const repeatCount = 5

func Repeat(input string) (repated string) {
	for i := 0; i < repeatCount; i++ {
		repated += input
	}
	return
}
