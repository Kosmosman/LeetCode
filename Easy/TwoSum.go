func twoSum(nums []int, target int) []int {
    var d = make(map[int]int)
    var check bool
    var result []int
    for index, value := range nums {
        _, check = d[target - value]
        if check {
            result = append(result, index, d[target - value])
            break
        } else {
            d[value] = index
        }
    }
    return result
}
