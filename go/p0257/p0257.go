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

package p0257

import (
	"fmt"

	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func binaryTreePaths(root *TreeNode) []string {
	ret := make([]string, 0)

	_path := func(root *TreeNode, prefix string, fn interface{}) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			ret = append(ret, prefix+fmt.Sprintf("%d", root.Val))
			return
		}

		_fn := fn.(func(*TreeNode, string, interface{}))
		prefix = prefix + fmt.Sprintf("%d->", root.Val)
		_fn(root.Left, prefix, fn)
		_fn(root.Right, prefix, fn)
	}

	_path(root, "", _path)

	return ret
}
