func pivotIndex(nums []int) int {
	var sumRight, sumLeft int
	res := -1
	for _, val := range nums {
		sumRight += val
	}
	for i, val := range nums {
		sumRight -= val
		if i > 0 {
			sumLeft += nums[i-1]
		}
		if sumRight == sumLeft {
			res = i
			break
		}
	}
	return res
}
