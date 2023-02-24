import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	var dict = make(map[[3]int]int)
	var size = len(nums) - 1
	var start, end = 0, size
	var tmpStart, tmpEnd int
	var res [][]int
	var tmp [3]int
	var tmp2 []int
	if nums[start]*nums[end] <= 0 {
		for (end - start) > 1 {
			for (end-start) > 1 && float64(nums[start]*(-1)/2) > float64(nums[end]) {
				start++
				continue
			}
			tmpStart = start
			tmpEnd = end
			for (tmpEnd-tmpStart) > 1 && (nums[tmpStart]*(-1) <= nums[tmpEnd]*2) {
				if binarySearch(nums[tmpStart+1:tmpEnd], -(nums[tmpStart] + nums[tmpEnd])) {
					tmp[0] = nums[tmpStart]
					tmp[1] = -(nums[tmpStart] + nums[tmpEnd])
					tmp[2] = nums[tmpEnd]
					dict[tmp] = 1
				}
				tmpEnd--
			}
			if nums[start]*(-1)/2 > nums[end] {
				end--
			} else {
				start++
			}
		}
	}
	for key, _ := range dict {
		tmp2 = append(tmp2, key[0], key[1], key[2])
		res = append(res, tmp2)
		tmp2 = nil
	}
	return res
}

func binarySearch(massive []int, desiredNumber int) bool {
	var mid, low, high = 0, 0, len(massive) - 1
	var flag bool
	for !flag && high >= low {
		mid = (low + high) / 2
		if massive[mid] == desiredNumber {
			flag = true
		} else if massive[mid] > desiredNumber {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return flag
}
