func isSubsequence(s string, t string) bool {
    var posSub = 0
    var res bool
    if len(s) > 0 {
        for _,val := range t {
            if val == rune(s[posSub]) {
                posSub++
            }
            if posSub == len(s) {
                res = true
                break
            }
        }
        fmt.Println(posSub)
        if posSub == len(s) {
            res = true
        }
    } else {
        res = true
    }
    return res
}
