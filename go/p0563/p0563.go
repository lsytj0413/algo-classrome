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

package p0563

import (
	"math"

	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func findTilt(root *TreeNode) int {
	_walk := func(n *TreeNode, fn interface{}) (tilt int, sum int) {
		if n == nil {
			return 0, 0
		}

		if n.Left == nil && n.Right == nil {
			return 0, n.Val
		}

		_fn := fn.(func(*TreeNode, interface{}) (int, int))
		tiltr, sumr := _fn(n.Right, fn)
		tiltl, suml := _fn(n.Left, fn)

		return int(math.Abs(float64(sumr-suml))) + tiltl + tiltr, n.Val + sumr + suml
	}

	tilt, _ := _walk(root, _walk)
	return tilt
}
