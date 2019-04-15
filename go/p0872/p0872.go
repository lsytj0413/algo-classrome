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

package p0872

import (
	"container/list"

	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	_leafs := func(n *TreeNode) []int {
		ret := make([]int, 0)
		if n == nil {
			return ret
		}
		stack := list.New()
		stack.PushBack(n)

		for stack.Len() != 0 {
			e := stack.Back()
			stack.Remove(e)

			node := e.Value.(*TreeNode)
			if node.Left == nil && node.Right == nil {
				ret = append(ret, node.Val)
			}

			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
		}

		return ret
	}

	leafs1 := _leafs(root1)
	leafs2 := _leafs(root2)

	if len(leafs2) != len(leafs1) {
		return false
	}

	for i := 0; i < len(leafs1); i++ {
		if leafs1[i] != leafs2[i] {
			return false
		}
	}

	return true
}
