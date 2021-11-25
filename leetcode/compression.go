package main

import "fmt"

func compress(chars []byte) int {

	if chars == nil {
		return 0
	} else if len(chars) == 1 {
		return 1
	} else {
		var i, j int = 0, 1
		for i < len(chars) && j < len(chars) {
			if chars[i] == chars[j] {
				j++
			} else {
				i++
				chars[i] = chars[j]
				j++
			}
		}
		return j
	}
}

func main() {
	var chars []byte = []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(chars))
}
