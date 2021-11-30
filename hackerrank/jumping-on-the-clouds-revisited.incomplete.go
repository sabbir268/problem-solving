package main

import "fmt"

func jumpingOnClouds(c []int32, k int32) int32 {
	var power int32 = 100
	var i int32 = 0 + k
	var n int32 = int32(len(c))
	for i < n {
		if c[i] == 1 {
			power = power - 1 - 2
		} else {
			power = power - 1
		}
		fmt.Println(i)
		fmt.Println(power)
		if n-i > k {
			i += k
		} else {
			i++
		}

	}
	return power
}

func main() {
	c := []int32{1, 1, 1, 0, 0, 1, 0, 0, 0, 0}
	k := int32(3)
	res := jumpingOnClouds(c, k)

	fmt.Println("res", res)
}
