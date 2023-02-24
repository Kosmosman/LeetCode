func isIsomorphic(s string, t string) bool {
	var dict = make(map[rune]rune)
    var dict2 = make(map[rune]rune)
	var res bool = true
	for i, val := range s {
		_, ok := dict[val]
		if ok {
			if dict[val] != rune(t[i]) {
				res = false
				break
			}
		} else {
			dict[val] = rune(t[i])
		}
        _, ok2 := dict2[rune(t[i])]
		if ok2 {
			if dict2[rune(t[i])] != val {
				res = false
				break
			}
		} else {
			dict2[rune(t[i])] = val
		}
	}
	return res
}
