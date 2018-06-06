package p0002

// ListNode for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	end := head
	carry := 0

	for l1 != nil || l2 != nil {
		if l1 != nil {
			carry = carry + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry = carry + l2.Val
			l2 = l2.Next
		}

		end.Next = &ListNode{
			Val:  carry % 10,
			Next: nil,
		}
		end = end.Next
		carry = carry / 10
	}

	if carry != 0 {
		end.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
		end = end.Next
	}

	return head.Next
}
