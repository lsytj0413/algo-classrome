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

package p0437

import (
	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	_sum := func(n *TreeNode, sum int, fn interface{}) int {
		_fn := fn.(func(*TreeNode, int, interface{}) int)
		if n == nil {
			return 0
		}

		c := 0
		if n.Val == sum {
			c++
		}

		return c + _fn(n.Left, sum-n.Val, fn) + _fn(n.Right, sum-n.Val, fn)
	}

	return _sum(root, sum, _sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}
