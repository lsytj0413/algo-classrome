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

package p0203

import (
	"algo/comm"
)

// ListNode for singly-linked list
type ListNode = comm.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{}
	tail := newHead

	for head != nil {
		if head.Val != val {
			tail.Next = head
			tail = tail.Next
		}
		head = head.Next
	}
	tail.Next = nil

	return newHead.Next
}
