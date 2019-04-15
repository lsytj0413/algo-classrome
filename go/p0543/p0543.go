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

package p0543

import (
	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	_walk := func(root *TreeNode, fn interface{}) (d int, l int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			return 0, 1
		}

		_fn := fn.(func(*TreeNode, interface{}) (int, int))
		dl, ll := _fn(root.Left, fn)
		dr, lr := _fn(root.Right, fn)

		l = ll + 1
		if ll < lr {
			l = lr + 1
		}

		d = ll + lr
		if d < dl {
			d = dl
		}
		if d < dr {
			d = dr
		}
		return
	}

	d, _ := _walk(root, _walk)
	return d
}
