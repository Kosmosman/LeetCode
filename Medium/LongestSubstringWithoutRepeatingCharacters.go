func lengthOfLongestSubstring(s string) int {
    dict := make(map[rune]int)
    var min_index, max_index, size, tmp int
    var check bool
    if len(s) > 0 {
        size = 1
    }
    for index, value := range s {
        _, check = dict[value]
        if check && dict[value] >= min_index {
            tmp = max_index - min_index
            if tmp > size {
                size = tmp
            }
            if (min_index > dict[value]) {
                min_index = index
            } else {
                min_index = dict[value] + 1
            }
            dict[value] = index
        } else {
            dict[value] = index
        }
        max_index++
    }
    if size > 0 && max_index - min_index > size {
        size = max_index - min_index
    }
    return size
}
