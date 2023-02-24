func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var rest int
	res := new(ListNode)
	out := res
	for l1 != nil || l2 != nil || rest != 0 {
		out.Val = (l1.Val + l2.Val + rest) % 10
		rest = (l1.Val + l2.Val) / 10
        if l1 != nil || l2 != nil || rest != 0 {
            out.Next = new(ListNode)
            out = out.Next
        }
        if l1.Next != nil {
            l1 = l1.Next
        } else {
            l1.Val = 0
        }
        if l2.Next != nil {
            l2 = l2.Next
        } else {
            l2.Val = 0
        }
	}
	return out
}
