package main

import "fmt"

func maiorNumeroRecursivo(nums []int, next int) int {

	if len(nums) == 1 {
		if next > nums[0] {
			return next
		}
		return nums[0]
	}

	if next < nums[0] {
		next = nums[0]
	}

	return maiorNumeroRecursivo(nums[1:], next)
}

func main() {

	nums := []int{122, 45, 2, 3, 4, 5, 67, 100}

	fmt.Println(maiorNumeroRecursivo(nums, 0))
}
