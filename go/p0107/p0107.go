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

package p0107

import (
	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func levelOrderBottom(root *TreeNode) [][]int {
	_levelOrderBottom := func(nodes []*TreeNode) ([]int, []*TreeNode) {
		nextNodes := make([]*TreeNode, 0)
		v := make([]int, 0)

		for _, n := range nodes {
			if n != nil {
				v = append(v, n.Val)
				nextNodes = append(nextNodes, n.Left, n.Right)
			}
		}

		return v, nextNodes
	}

	ret := make([][]int, 0)
	nextNodes := make([]*TreeNode, 1)
	nextNodes[0] = root
	v := make([]int, 0)
	for len(nextNodes) != 0 {
		v, nextNodes = _levelOrderBottom(nextNodes)
		if len(v) != 0 {
			ret = append(ret, v)
		}
	}

	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}

	return ret
}
