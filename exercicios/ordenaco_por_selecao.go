package main

import "fmt"

func menor_valor(nums []int) (index int) {
	menor := nums[0]

	for i, v := range nums {
		if v < menor {
			index = i
			menor = v
		}
	}

	return
}

func ordenar(nums []int) (newNums []int) {

	l := len(nums)

	for i := 0; i < l; i++ {
		menor := menor_valor(nums)
		newNums = append(newNums, nums[menor])
		nums = append(nums[:menor], nums[menor+1:]...)
	}

	return
}

func main() {

	nums := []int{-2, 3, 4, 5, -1, 6, 7, 8, 9}

	// index := busca_menor_valor(nums)
	// fmt.Println(index)
	// fmt.Println(nums)

	fmt.Println(ordenar(nums))
}
