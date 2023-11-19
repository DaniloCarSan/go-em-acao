package main

import (
	"fmt"
)

func qsort(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}

	pivo := nums[0]

	menores := []int{}
	maiores := []int{}

	for _, v := range nums[1:] {
		if v <= pivo {
			menores = append(menores, v)
		}
		if v > pivo {
			maiores = append(maiores, v)
		}
	}

	r1 := qsort(menores)
	r2 := qsort(maiores)

	r1 = append(r1, pivo)
	r1 = append(r1, r2...)
	return r1
}

func main() {

	nums := []int{10, 0, 1, 8}

	fmt.Println(qsort(nums))
}
