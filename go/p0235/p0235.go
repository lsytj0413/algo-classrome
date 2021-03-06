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

package p0235

import "github.com/lsytj0413/algo-classrome/go/comm"

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}

	if root.Val >= p.Val && root.Val <= q.Val {
		return root
	}

	if root.Val < p.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return lowestCommonAncestor(root.Left, p, q)
}
