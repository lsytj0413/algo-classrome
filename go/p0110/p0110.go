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

package p0110

import (
	"algo/comm"
)

// TreeNode ...
type TreeNode = comm.TreeNode

func isBalanced(root *TreeNode) bool {
	_depth := func(n *TreeNode, fn interface{}) (int, bool) {
		if n == nil {
			return 0, true
		}

		_fn := fn.(func(*TreeNode, interface{}) (int, bool))
		left, isBalance := _fn(n.Left, fn)
		if !isBalance {
			return 0, false
		}

		right, isBalance := _fn(n.Right, fn)
		if !isBalance {
			return 0, false
		}

		if left > right+1 || right > left+1 {
			return 0, false
		}

		if right > left {
			return right + 1, true
		}
		return left + 1, true
	}

	_, ret := _depth(root, _depth)
	return ret
}
