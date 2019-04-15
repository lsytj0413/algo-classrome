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

package p0501

import (
	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func findMode(root *TreeNode) []int {
	ret, maxCount := make([]int, 0), 0
	last, lastCount := 0, -1

	_walk := func(n *TreeNode, fn interface{}) {
		if n == nil {
			return
		}

		_fn := fn.(func(*TreeNode, interface{}))
		if n.Left != nil {
			_fn(n.Left, fn)
		}

		if lastCount == -1 {
			last, lastCount = n.Val, 1
		} else {
			if n.Val == last {
				lastCount++
			} else {
				if lastCount == maxCount {
					ret = append(ret, last)
				} else if lastCount > maxCount {
					ret, maxCount = []int{last}, lastCount
				}

				last, lastCount = n.Val, 1
			}
		}

		if n.Right != nil {
			_fn(n.Right, fn)
		}
	}

	_walk(root, _walk)

	if lastCount == maxCount {
		ret = append(ret, last)
	} else if lastCount > maxCount {
		ret, maxCount = []int{last}, lastCount
	}

	return ret
}
