package main

import "fmt"

func sum(nums []int) int {
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}

func sumOfAllSubarrays(nums []int, k int) []int {
	currentSubArraySum := sum(nums[:k])
	result := []int{currentSubArraySum}
	for i := 1; i < len(nums) - k + 1; i++ {
		currentSubArraySum = currentSubArraySum - nums[i - 1]
		currentSubArraySum = currentSubArraySum + nums[i + k - 1]
		result = append(result, currentSubArraySum)
	}
	return result
}

func main() {
	fmt.Println(sumOfAllSubarrays([]int{1, 2, 3, 4, 5}, 4))
}
