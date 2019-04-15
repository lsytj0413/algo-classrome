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

package p0783

import (
	"container/list"
	"math"

	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return 0
	}

	pre := root
	stack, min, last := list.New(), math.MaxInt32, math.MaxInt32

	for stack.Len() != 0 || pre != nil {
		if pre != nil {
			stack.PushBack(pre)
			pre = pre.Left
		} else {
			e := stack.Back()
			stack.Remove(e)

			pre = e.Value.(*TreeNode)

			if last != math.MaxInt32 && pre.Val-last < min {
				min = pre.Val - last
			}
			last = pre.Val

			pre = pre.Right
		}
	}

	return min
}
