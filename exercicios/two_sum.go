package main

import "fmt"

func twoSum(nums []int, target int) []int {

	subs := map[int]int{}

	if len(subs) == 2 {
		return []int{0, 1}
	}

	for i, num := range nums {

		if k, ok := subs[num]; ok && k != i {
			return []int{k, i}
		}

		subs[target-num] = i
	}

	return []int{}
}

func main() {

	nums := []int{11, 1, 2, 3, 4, 5, 7, 8, 9, 10, 11, 12}

	fmt.Println(twoSum(nums, 11))

}
