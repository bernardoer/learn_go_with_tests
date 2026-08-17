package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	// Function create on class->

	// var repeated strings.Builder
	// for i := 0; i < repeatCount; i++ {
	// 	repeated.WriteString(character)
	// }

	// Use String library function to do the same job
	return strings.Repeat(character, repeatCount)
}
