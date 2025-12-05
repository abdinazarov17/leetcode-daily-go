package main

import "fmt"

/*
	link - https://leetcode.com/problems/two-sum/
	O(n)
*/

func twoSum(nums []int, target int) []int {

	numsMap := make(map[int]int)
	for i, v := range nums {
		numsMap[v] = i + 1
	}

	for i, v := range nums {
		if numsMap[target-v] > 0 && i != numsMap[target-v]-1 {
			return []int{i, numsMap[target-v] - 1}
		}
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
