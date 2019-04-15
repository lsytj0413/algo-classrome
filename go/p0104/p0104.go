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

package p0104

import (
	"math"

	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMax := 1 + maxDepth(root.Left)
	rightMax := 1 + maxDepth(root.Right)
	return int(math.Max(float64(leftMax), float64(rightMax)))
}
