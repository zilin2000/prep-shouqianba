package iteration

const repeatCount = 5

func Repeat(input string) (output string) {
	for i := 0; i < repeatCount; i++ {
		output += input
	}
	return
}
