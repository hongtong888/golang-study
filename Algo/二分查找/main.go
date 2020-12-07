package main

import "fmt"

/**
二分查找的思路是，先拿排好序数列的中位数与目标数字 189 对比，如果刚好匹配目标，结束。
如果中位数比目标数字大，因为已经排好序，所以中位数右边的数字绝对都比目标数字大，那么从中位数的左边找。
如果中位数比目标数字小，因为已经排好序，所以中位数左边的数字绝对都比目标数字小，那么从中位数的右边找。
这种分而治之，一分为二的查找叫做二分查找算法。
*/
func main() {
	array := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
	target := 81
	result := BinarySearch1(array, target, 0, len(array)-1)
	fmt.Println(target, result)
}

/**
采用递归的思想
*/
func BinarySearch(array []int, target int, start, end int) int {
	if start > end {
		//出界了找不到
		return -1
	}
	//算出中位数，从中间开始找
	mid := (start + end) / 2
	if target == array[mid] {
		return mid
	}
	//如果目标数小于中位数，从左边开始，如果目标数大于中位数，从右边开始
	if target > array[mid] {
		return BinarySearch(array, target, mid, end)
	} else {
		return BinarySearch(array, target, start, mid)
	}
}

/**
采用普通方法
*/
func BinarySearch1(array []int, target int, start, end int) int {
	startTemp := start
	endTemp := end
	for {
		if startTemp > endTemp {
			//出界了找不到
			return -1
		}
		//算出中位数，从中间开始找
		mid := (startTemp + endTemp) / 2
		if target == array[mid] {
			return mid
		} else if target < array[mid] {
			endTemp = mid - 1
		} else {
			startTemp = mid + 1
		}
	}
}
