/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
*/

package eetcode

func TwoSum(nums []int, target int) []int {
	lens := len(nums)
	for index := 0; index < lens - 1; index++ {
		for jindex := index + 1; jindex < lens; jindex++ {
			if nums[index] + nums[jindex] == target {
				return []int{index, jindex}
			}
		}
	}
	return []int{}
}

