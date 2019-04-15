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

package p0530

import (
	"math"

	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func getMinimumDifference(root *TreeNode) int {
	last, minDiff := math.MinInt32, math.MaxInt32

	_walk := func(n *TreeNode, fn interface{}) {
		if n == nil {
			return
		}

		_fn := fn.(func(*TreeNode, interface{}))
		if n.Left != nil {
			_fn(n.Left, fn)
		}

		nowDiff := n.Val - last
		if nowDiff < minDiff {
			minDiff = nowDiff
		}
		last = n.Val

		if n.Right != nil {
			_fn(n.Right, fn)
		}
	}

	_walk(root, _walk)

	return minDiff
}
