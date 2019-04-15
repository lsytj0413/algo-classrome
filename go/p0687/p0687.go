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

package p0687

import (
	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func longestUnivaluePath(root *TreeNode) int {
	max := 0

	_max := func(v1 int, v2 int) int {
		if v1 > v2 {
			return v1
		}

		return v2
	}
	_walk := func(n *TreeNode, fn interface{}) int {
		if n == nil {
			return 0
		}

		left, right := 0, 0
		_fn := fn.(func(*TreeNode, interface{}) int)
		if n.Left != nil {
			left = _fn(n.Left, fn)
		}
		if n.Right != nil {
			right = _fn(n.Right, fn)
		}

		if n.Left != nil && n.Left.Val == n.Val {
			left++
		} else {
			left = 0
		}
		if n.Right != nil && n.Right.Val == n.Val {
			right++
		} else {
			right = 0
		}
		if (left > 0 || right > 0) && left+right > max {
			max = left + right
		}
		return _max(left, right)
	}
	_walk(root, _walk)

	return max
}
