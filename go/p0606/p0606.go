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

package p0606

import (
	"fmt"

	"algo/comm"
)

// TreeNode is binary tree node
type TreeNode = comm.TreeNode

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	switch {
	case t.Left == nil && t.Right == nil:
		return fmt.Sprintf("%d", t.Val)
	case t.Left == nil:
		return fmt.Sprintf("%d()(%s)", t.Val, tree2str(t.Right))
	case t.Right == nil:
		return fmt.Sprintf("%d(%s)", t.Val, tree2str(t.Left))
	default:
		return fmt.Sprintf("%d(%s)(%s)", t.Val, tree2str(t.Left), tree2str(t.Right))
	}
}
