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

package p0474

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	_max := func(v1 int, v2 int) int {
		if v1 > v2 {
			return v1
		}
		return v2
	}

	for _, str := range strs {
		zeros, ones := 0, 0
		for i := 0; i < len(str); i++ {
			if str[i] == '1' {
				ones++
			} else {
				zeros++
			}
		}

		for i := m; i > zeros-1; i-- {
			for j := n; j > ones-1; j-- {
				dp[i][j] = _max(dp[i][j], dp[i-zeros][j-ones]+1)
			}
		}
	}
	return dp[m][n]
}
