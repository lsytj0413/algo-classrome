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

package p9999

// 有一个N行M列的棋盘格, 有个国际象棋里的马要从这个棋盘格的第一行跳到最后一行。要求这匹马只能从上往下跳，不能倒着跳，即只能跳往下一行或者下面第二行。
// 每个格子上有一个数字，请为小马寻找一条路径，要求路径上所有数字之和最小。小马可以从第一行的任意某个格子开始，也必须到最后一行的某个格子结束。

// 输入：一个NxM的矩阵
// 输出：一个数字，这条路径上所有数之和

// 示例：

// 输入：
// | 3 | 0 | -2 | 4 | 0 |
// | -1 | 2 | -2 | 1 | 4
// | 3 | 1 | -2 | -3 | 3 |
// | 2 | -4 | -3 | -3 | 2 |
// | 5 | 2 | -2 | -3 | 1 |

// 输出：-12
// 说明：跳的顺序为 (0,2) -> (2, 3) -> (3, 1) -> (4, 3)

// 附加题：
// 如果允许小马往回跳，算法有什么不同？

func minValue(grid [][]int) int {
	_min := func(arr []int) int {
		v := arr[0]
		for i := range arr {
			if arr[i] < v {
				v = arr[i]
			}
		}

		return v
	}

	row, col := len(grid), len(grid[0])
	for i := 1; i < row; i++ {
		for j := 0; j < col; j++ {
			v := make([]int, 0, 4)
			if i >= 2 {
				if j > 0 {
					v = append(v, grid[i-2][j-1])
				}
				if j < col-1 {
					v = append(v, grid[i-2][j+1])
				}
			}

			if j > 1 {
				v = append(v, grid[i-1][j-2])
			}
			if j < col-2 {
				v = append(v, grid[i-1][j+2])
			}

			grid[i][j] = grid[i][j] + _min(v)
		}
	}

	return _min(grid[len(grid)-1])
}
