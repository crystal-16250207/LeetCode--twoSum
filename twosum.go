package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
func main() {
	nums := []int{3, 2, 4}
	target := 6
	res := twoSum(nums, target)
	fmt.Println(res)
}
