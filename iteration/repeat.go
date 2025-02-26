package iteration

import "strings"

// func Repeat(character string) string {
// 	var repeated string

// 	for i := 0; i < 5; i++ {
// 		repeated += character
// 	}

// 	return repeated
// }

// const repeatCount = 5

func Repeat(character string, counter int) string {
	var repeated strings.Builder

	for i := 0; i < counter; i++ {
		repeated.WriteString(character)
	}

	return repeated.String()
}
