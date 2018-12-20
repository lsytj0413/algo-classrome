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

package p0375

import (
	"math"
)

func getMoneyAmount(n int) int {
	if 1 == n {
		return 0
	}

	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	_max := func(x int, y int) int {
		if x > y {
			return x
		}
		return y
	}
	_min := func(x int, y int) int {
		if x > y {
			return y
		}
		return x
	}

	for i := 2; i <= n; i++ {
		for j := i - 1; j > 0; j-- {
			min := math.MaxInt32
			for k := j + 1; k < i; k++ {
				min = _min(min, k+_max(dp[j][k-1], dp[k+1][i]))
			}

			if j+1 == i {
				dp[j][i] = j
			} else {
				dp[j][i] = min
			}
		}
	}

	return dp[1][n]
}
