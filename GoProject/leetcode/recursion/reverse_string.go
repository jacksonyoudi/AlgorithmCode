package main

import "fmt"

func reverseString(s []byte) []byte {
	if len(s) <= 1 {
		return s
	}
	result := reverseString(s[1:])
	result = append(result, s[0])
	return result
}


func main() {
	a := "abcdef"
	b := reverseString([]byte(a))
	fmt.Println(string(b))

}