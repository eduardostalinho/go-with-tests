package iteration

// Repeat returns a string repeated 5 times
func Repeat(c string, t int) string {
	var r string
	for i := 0; i < t; i++ {
		r = r + c
	}
	return r
}
