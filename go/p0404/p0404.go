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

package p0404

import (
	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0

	_sum := func(root *TreeNode, isLeft bool, fn interface{}) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			if isLeft {
				sum += root.Val
			}
			return
		}

		_fn := fn.(func(*TreeNode, bool, interface{}))
		_fn(root.Left, true, fn)
		_fn(root.Right, false, fn)
	}

	_sum(root, false, _sum)

	return sum
}
