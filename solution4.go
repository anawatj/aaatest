package main

import "fmt"

func solution4(a []int, k int) []int {
	var ret []int

	if len(a) > 100 {
		return ret
	}

	for index := 0; index < len(a); index++ {
		var currentIndex int

		if index+k > len(a)-1 {
			fmt.Println(index + k)
			currentIndex = ((index + k) % len(a))
			item := a[index]
			ret[currentIndex] = item

		} else {
			fmt.Println(index + k)
			currentIndex = index + k
			item := a[index]
			ret[currentIndex] = item

		}

	}
	return ret

}
