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

package p0637

import (
	"container/list"

	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func averageOfLevels(root *TreeNode) []float64 {
	ret := make([]float64, 0)
	if root == nil {
		return ret
	}

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() != 0 {
		count, sum, i := stack.Len(), 0, 0
		for i < count {
			e := stack.Front()
			stack.Remove(e)

			n := e.Value.(*TreeNode)
			sum += n.Val
			if n.Left != nil {
				stack.PushBack(n.Left)
			}
			if n.Right != nil {
				stack.PushBack(n.Right)
			}

			i++
		}

		ret = append(ret, float64(sum)/float64(count))
	}

	return ret
}
