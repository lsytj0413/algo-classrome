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

package p0108

import (
	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	if 0 == len(nums) {
		return nil
	}

	if 1 == len(nums) {
		return &TreeNode{
			Val: nums[0],
		}
	}

	mid := len(nums) / 2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  nil,
		Right: nil,
	}

	root.Left = sortedArrayToBST(nums[0 : len(nums)/2])
	root.Right = sortedArrayToBST(nums[len(nums)/2+1 : len(nums)])

	return root
}
