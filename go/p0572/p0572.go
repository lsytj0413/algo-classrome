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

package p0572

import (
	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func isSubtree(s *TreeNode, t *TreeNode) bool {
	_isSameTree := func(s *TreeNode, t *TreeNode, cmp interface{}) bool {
		if s == nil && t == nil {
			return true
		}

		if s == nil || t == nil {
			return false
		}

		_cmp := cmp.(func(*TreeNode, *TreeNode, interface{}) bool)
		return s.Val == t.Val && _cmp(s.Left, t.Left, cmp) && _cmp(s.Right, t.Right, cmp)
	}

	_walk := func(s *TreeNode, t *TreeNode, fn interface{}) bool {
		_fn := fn.(func(*TreeNode, *TreeNode, interface{}) bool)

		if s == nil && t == nil {
			return true
		}

		if s == nil || t == nil {
			return false
		}

		if s.Val == t.Val {
			if _isSameTree(s, t, _isSameTree) {
				return true
			}
		}

		return _fn(s.Left, t, fn) || _fn(s.Right, t, fn)
	}

	return _walk(s, t, _walk)
}
