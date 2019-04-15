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

package p0653

import (
	"container/list"

	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func findTarget(root *TreeNode, k int) bool {
	_walk := func(n *TreeNode) bool {
		maps := make(map[int]int)
		if n == nil {
			return false
		}

		stack := list.New()
		stack.PushBack(n)

		for stack.Len() != 0 {
			e := stack.Back()
			stack.Remove(e)

			node := e.Value.(*TreeNode)

			if _, exists := maps[k-node.Val]; exists {
				return true
			}
			maps[node.Val] = 1

			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
		}

		return false
	}

	return _walk(root)
}
