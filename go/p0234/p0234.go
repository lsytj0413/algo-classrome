// Copyright (c) 2018 soren yang
//
// Licensed under the MIT License
// you may not use this file except in complicance with the License.
// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0234

import (
	"algo/comm"
)

// ListNode for singly-linked list
type ListNode = comm.ListNode

func isPalindrome(head *ListNode) bool {
	_len := func(n *ListNode) int {
		l := 0
		for n != nil {
			l++
			n = n.Next
		}

		return l
	}

	l := _len(head)
	if l == 0 {
		return true
	}

	// 反转前半部分链表
	var nhead *ListNode
	for i := 0; i < l/2; i++ {
		t := head.Next
		head.Next = nhead
		nhead = head
		head = t
	}

	if l%2 == 1 {
		head = head.Next
	}

	for head != nil {
		if nhead.Val != head.Val {
			return false
		}

		head = head.Next
		nhead = nhead.Next
	}

	return true
}
