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

package p0378

// TODO: 优先级队列
func kthSmallest(matrix [][]int, k int) int {
	_helper := func(matrix [][]int, n int) int {
		i, j := len(matrix)-1, 0
		r := 0
		for i >= 0 && j < len(matrix[0]) {
			if matrix[i][j] > n {
				i--
			} else {
				r = r + i + 1
				j++
			}
		}
		return r
	}

	l := len(matrix)
	low, high := matrix[0][0], matrix[l-1][l-1]
	for low <= high {
		mid := low + (high-low)/2
		if _helper(matrix, mid) < k {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
