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

package p0024

import (
	"github.com/lsytj0413/algo-classrome/go/comm"
)

// ListNode for singly-linked list
type ListNode = comm.ListNode

func swapPairs(head *ListNode) *ListNode {
	n := &ListNode{
		Val:  0,
		Next: head,
	}
	pre := n
	for pre.Next != nil && pre.Next.Next != nil {
		a := pre.Next
		b := a.Next
		pre.Next, b.Next, a.Next = b, a, b.Next
		pre = a
	}

	return n.Next
}
