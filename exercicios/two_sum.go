package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var index1, index2 int

	for j, num1 := range nums {
		for k, num2 := range nums {
			if r := num1 + num2; r == target {
				if j != k {
					index1 = j
					index2 = k
					break
				}

			}
		}
	}

	return []int{index1, index2}
}

func main() {

	nums := []int{11, 1, 2, 3, 4, 5, 7, 8, 9, 10, 11, 12}

	fmt.Println(twoSum(nums, 11))

}
