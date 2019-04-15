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

package p0897

import (
	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	nroot := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	tail := nroot

	_walk := func(root *TreeNode, fn interface{}) {
		if root == nil {
			return
		}

		_fn := fn.(func(*TreeNode, interface{}))
		if root.Left != nil {
			_fn(root.Left, fn)
		}

		tail.Right = &TreeNode{
			Val:   root.Val,
			Left:  nil,
			Right: nil,
		}
		tail = tail.Right

		if root.Right != nil {
			_fn(root.Right, fn)
		}
	}

	_walk(root, _walk)

	return nroot.Right
}
