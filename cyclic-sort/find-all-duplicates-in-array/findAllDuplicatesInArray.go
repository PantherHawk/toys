package main

import "fmt"

func findAllDupsInArray(nums []int) []int {
	i := 0
	n := len(nums)
	output := []int{}
	
	for i < n {
		fmt.Println(nums)
		fmt.Println(i)
		fmt.Println(nums[i])
		correct := nums[i] - 1
		if nums[i] != nums[correct] {
			nums[i], nums[correct] = nums[correct], nums[i]
		} else {
			i++
		}
	}
	fmt.Println(nums)
	for i, v := range(nums) {
		if i + 1 != v {
			output = append(output, v)
		}
	}
	return output
}

func main() {
	fmt.Println(findAllDupsInArray([]int{4, 3, 2, 7,  8, 2, 3, 1}))
}
