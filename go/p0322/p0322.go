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

package p0322

import (
	"math"
	"sort"
)

func coinChange(coins []int, amount int) int {
	// DOC: 遍历所有可能, timeout
	_f1 := func(coins []int, amount int) int {
		sort.Ints(coins)
		min := math.MaxInt32

		_helper := func(coins []int, amount int, count int, fn interface{}) {
			if amount == 0 {
				if count < min {
					min = count
				}
			}

			_fn := fn.(func([]int, int, int, interface{}))
			for i := len(coins) - 1; i >= 0; i-- {
				if amount >= coins[i] {
					_fn(coins, amount-coins[i], count+1, fn)
				}
			}
		}
		_helper(coins, amount, 0, _helper)

		if min == math.MaxInt32 {
			return -1
		}
		return min
	}
	_ = _f1

	_f2 := func(coins []int, amount int) int {
		dp := make([]int, amount+1)
		dp[0] = 0

		for i := 1; i <= amount; i++ {
			min := math.MaxInt32
			for j := 0; j < len(coins); j++ {
				if i < coins[j] || dp[i-coins[j]] == -1 {
					continue
				}

				count := dp[i-coins[j]] + 1
				if count < min {
					min = count
				}
			}

			if min == math.MaxInt32 {
				dp[i] = -1
			} else {
				dp[i] = min
			}
		}

		return dp[amount]
	}
	_ = _f2

	return _f2(coins, amount)
}
