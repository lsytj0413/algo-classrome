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

package p0002

import (
	"algo/comm"
)

// ListNode for singly-linked list
type ListNode = comm.ListNode

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
