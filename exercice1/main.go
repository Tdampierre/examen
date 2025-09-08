package main

func main() {
	println(Surprise("ABCDabc"))
	println(surprise_annexe(1))
}

func Surprise(s string) bool {
	if len(s) < 1 {
		return false
	}
	return surprise_annexe(s[0]) || Surprise(s[1:])
}
func surprise_annexe(a byte) bool {
	return (a <= 'z' && a >= 'a') || (a <= 'Z' && a >= 'A')
}
