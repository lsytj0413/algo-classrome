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

package p0111

import (
	"math"

	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left != nil && root.Right != nil {
		leftMin := 1 + minDepth(root.Left)
		rightMin := 1 + minDepth(root.Right)

		return int(math.Min(float64(leftMin), float64(rightMin)))
	}

	if root.Left != nil {
		return 1 + minDepth(root.Left)
	}

	return 1 + minDepth(root.Right)
}
