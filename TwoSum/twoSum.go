package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		m[val] = index
	}

	for index, val := range nums {
		compliment := target - val
		if j, exists := m[compliment]; exists && j != index {
			return []int{index, j}
		}
	}

	return nil
}

func main() {
	nums := []int{3, 3}
	target := 6
	fmt.Printf("%+v", twoSum(nums, target))
}
