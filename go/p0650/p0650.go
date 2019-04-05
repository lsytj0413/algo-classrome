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

package p0650

// TODO: 因式分解
func minSteps(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = i
	}

	_min := func(x int, y int) int {
		if x < y {
			return x
		}
		return y
	}
	for i := 6; i <= n; i++ {
		for j := i / 2; j >= 2; j-- {
			if i%j == 0 {
				dp[i] = _min(dp[i], dp[j]+i/j)
			}
		}
	}

	return dp[n]
}
