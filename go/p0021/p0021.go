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

package p0021

import (
	"algo/comm"
)

// ListNode for singly-linked list
type ListNode = comm.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	end := head

	for l1 != nil || l2 != nil {
		if l1 == nil {
			end.Next = l2
			l2 = l2.Next
		} else if l2 == nil {
			end.Next = l1
			l1 = l1.Next
		} else {
			if l1.Val < l2.Val {
				end.Next = l1
				l1 = l1.Next
			} else {
				end.Next = l2
				l2 = l2.Next
			}
		}
		end = end.Next
	}
	return head.Next
}
