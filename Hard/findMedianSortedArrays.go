func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var index_nums1, index_nums2 int
	var max, result_1, result_2 float64
	var first_index, second_index, last_index int = 0, -1, -1
	result_1 = -1
	result_2 = -1
	if (len(nums1)+len(nums2))%2 == 0 {
		first_index = (len(nums1)+len(nums2))/2 - 1
		second_index = first_index + 1
		last_index = second_index
	} else {
		first_index = (len(nums1) + len(nums2)) / 2
		last_index = first_index
	}
	for (index_nums1 < len(nums1) && index_nums1 <= last_index) || (index_nums2 < len(nums2) && index_nums2 <= last_index) {
		if (index_nums1 >= len(nums1)) || (index_nums2 < len(nums2) && nums1[index_nums1] > nums2[index_nums2]) {
			max = float64(nums2[index_nums2])
			if result_1 < 0 {
				result_1 = addFirst(index_nums1, index_nums2, first_index, max)
			}
			if second_index >= 0 && result_2 < 0 {
				result_2 = addFirst(index_nums1, index_nums2, second_index, max)
			}
			index_nums2++
		} else if index_nums1 < len(nums1) {
			max = float64(nums1[index_nums1])
			if result_1 < 0 {
				result_1 = addFirst(index_nums1, index_nums2, first_index, max)
			}
			if second_index >= 0 && result_2 < 0 {
				result_2 = addFirst(index_nums1, index_nums2, second_index, max)
			}
			index_nums1++
		}
	}
	if second_index >= 0 {
		max = (result_1 + result_2) / 2
	} else {
		max = result_1
	}
	if max < 0 && (len(nums1)+len(nums2) == 0) {
		max = 0
	}
	return max
}

func addFirst(index_1, index_2, found_index int, max float64) float64 {
	var result float64 = -1
	if (index_1 + index_2) == found_index {
		result = max
	}
	return result
}
