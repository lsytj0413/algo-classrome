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

package p0671

import (
	"container/list"
	"math"

	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func findSecondMinimumValue(root *TreeNode) int {
	first, second := math.MaxInt32, math.MaxInt32

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() != 0 {
		e := stack.Front()
		stack.Remove(e)

		n := e.Value.(*TreeNode)

		if n.Left != nil {
			stack.PushBack(n.Left)
		}
		if n.Right != nil {
			stack.PushBack(n.Right)
		}

		switch {
		case n.Val < first:
			first, second = n.Val, first
		case n.Val > first && n.Val < second:
			second = n.Val
		}
	}

	if second == math.MaxInt32 {
		return -1
	}
	return second
}
