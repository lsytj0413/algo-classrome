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

package p0343

import (
	"math"
)

// TODO: 公式解法
func integerBreak(n int) int {
	_f1 := func(n int) int {
		dp := make([]int, n+1)
		dp[2] = 1

		for i := 3; i <= n; i++ {
			max := 0
			for j := 2; j < i; j++ {
				if j*(i-j) > max {
					max = j * (i - j)
				}
				if dp[j]*(i-j) > max {
					max = dp[j] * (i - j)
				}
			}
			dp[i] = max
		}

		return dp[n]
	}
	_ = _f1

	_f2 := func(n int) int {
		switch {
		case n == 2:
			return 1
		case n == 3:
			return 2
		case n%3 == 0:
			return int(math.Pow(float64(3), float64(n/3)))
		case n%3 == 1:
			return 2 * 2 * int(math.Pow(float64(3), float64((n-4)/3)))
		default:
			return 2 * int(math.Pow(float64(3), float64((n-2)/3)))
		}
	}
	_ = _f2

	return _f2(n)
}
