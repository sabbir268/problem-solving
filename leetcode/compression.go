package main

import "fmt"

func compress(chars []byte) int {

	if chars == nil {
		return 0
	} else if len(chars) == 1 {
		return 1
	} else {
		var i, j, count int
		for i = 0; i < len(chars); i++ {
			if i == 0 {
				count = 1
			} else if chars[i] != chars[i-1] {
				chars[j] = chars[i-1]
				j++
				if count > 1 {
					for _, c := range fmt.Sprintf("%d", count) {
						chars[j] = byte(c)
						j++
					}
				}
				count = 1
			} else {
				count++
			}
		}
		chars[j] = chars[i-1]
		j++
		if count > 1 {
			for _, c := range fmt.Sprintf("%d", count) {
				chars[j] = byte(c)
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
