package iteration

func Repeat(c string, x int) string {
	var repeated string
	for i := 0; i < x; i++ {
		repeated += c
	}
	return repeated
}
