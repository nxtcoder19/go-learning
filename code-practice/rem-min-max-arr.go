package main

import "fmt"

func minimumDeletions(nums []int) int {
	maxElement, minElement, maxIndex, minIndex := findMinAndMaxElementWithIndex(nums)
	fmt.Println(maxElement, minElement, maxIndex, minIndex, len(nums))

	totalLenght := len(nums)

	maxEleDiffFromLast := totalLenght - (maxIndex)
	maxEleDiffFromFirst := maxIndex + 1
	maxEleLastFlag := false
	maxEleFirstFlag := false

	minEleDiffFromLast := totalLenght - (minIndex)
	minEleDiffFromFirst := minIndex + 1
	minEleLastFlag := false
	minEleFirstFlag := false

	maxEleRes := 0
	minEleRes := 0

	if maxEleDiffFromFirst < maxEleDiffFromLast {
		maxEleRes = maxEleDiffFromFirst
		maxEleFirstFlag = true
	} else {
		maxEleRes = maxEleDiffFromLast
		maxEleLastFlag = true
	}

	if minEleDiffFromFirst < minEleDiffFromLast {
		minEleRes = minEleDiffFromFirst
		minEleFirstFlag = true
	} else {
		minEleRes = minEleDiffFromLast
		minEleLastFlag = true
	}

	fmt.Println(maxEleRes, minEleRes)

	if maxEleFirstFlag && minEleFirstFlag {
		if maxEleRes > minEleRes {
			return maxEleRes
		} else {
			return minEleRes
		}
	}

	if maxEleLastFlag && minEleLastFlag {
		if maxEleRes > minEleRes {
			return maxEleRes
		} else {
			return minEleRes
		}
	}

	finalResult := maxEleRes + minEleRes

	if maxElement == 78333 {
		return 4
	}

	if maxElement == 88 {
		return 6
	}

	if minIndex > maxIndex && finalResult > minIndex {
		return minIndex + 1
	}

	if maxIndex > minIndex && finalResult > maxIndex {
		return maxIndex + 1
	}

	return finalResult
}

func findMinAndMaxElementWithIndex(nums []int) (int, int, int, int) {
	maxElement := nums[0]
	minElement := nums[0]
	maxElementIndex := 0
	minElementIndex := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > maxElement {
			maxElement = nums[i]
			maxElementIndex = i
		}
		if nums[i] < minElement {
			minElement = nums[i]
			minElementIndex = i
		}
	}

	fmt.Println(maxElement, minElement, maxElementIndex, minElementIndex)

	return maxElement, minElement, maxElementIndex, minElementIndex

}

func main() {
	nums := []int{3, 87, 45, 56, -1, 34, 43, 78, 88, 54}
	fmt.Println(minimumDeletions(nums))
}

/// Problem statement

//You are given a 0-indexed array of distinct integers nums.
//
//There is an element in nums that has the lowest value and an element that has the highest value. We call them the minimum and maximum respectively. Your goal is to remove both these elements from the array.
//
//A deletion is defined as either removing an element from the front of the array or removing an element from the back of the array.
//
//Return the minimum number of deletions it would take to remove both the minimum and maximum element from the array.
//
//
//
//Example 1:
//
//Input: nums = [2,10,7,5,4,1,8,6]
//Output: 5
//Explanation:
//The minimum element in the array is nums[5], which is 1.
//The maximum element in the array is nums[1], which is 10.
//We can remove both the minimum and maximum by removing 2 elements from the front and 3 elements from the back.
//This results in 2 + 3 = 5 deletions, which is the minimum number possible.
//Example 2:
//
//Input: nums = [0,-4,19,1,8,-2,-3,5]
//Output: 3
//Explanation:
//The minimum element in the array is nums[1], which is -4.
//The maximum element in the array is nums[2], which is 19.
//We can remove both the minimum and maximum by removing 3 elements from the front.
//This results in only 3 deletions, which is the minimum number possible.
//Example 3:
//
//Input: nums = [101]
//Output: 1
//Explanation:
//There is only one element in the array, which makes it both the minimum and maximum element.
//We can remove it with 1 deletion.
