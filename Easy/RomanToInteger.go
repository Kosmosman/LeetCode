func romanToInt(s string) int {
	var res int
	var end = len(s) - 1
	for index, val := range s {
		if index < end {
			res = chooseSim(val, index, res, s)
		}
	}
	res = chooseSim(rune(s[end]), end-1, res, s)
	return res
}

func chooseSim(val rune, index, res int, s string) int {
	switch val {
	case 'I':
		if s[index+1] == 'X' || s[index+1] == 'V' {
			res -= 1
		} else {
			res += 1
		}
	case 'V':
		res += 5
	case 'X':
		if s[index+1] == 'L' || s[index+1] == 'C' {
			res -= 10
		} else {
			res += 10
		}
	case 'L':
		res += 50
	case 'C':
		if s[index+1] == 'D' || s[index+1] == 'M' {
			res -= 100
		} else {
			res += 100
		}
	case 'D':
		res += 500
	case 'M':
		res += 1000
	}
	return res
}
