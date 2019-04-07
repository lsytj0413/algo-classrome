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

package p0714

func maxProfit(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}

	dp1, dp2 := make([]int, len(prices)), make([]int, len(prices))
	dp1[0] = 0 - prices[0]

	_max := func(x int, y int) int {
		if x < y {
			return y
		}
		return x
	}
	for i := 1; i < len(prices); i++ {
		dp1[i] = _max(dp1[i-1], dp2[i-1]-prices[i])
		dp2[i] = _max(dp2[i-1], dp1[i-1]+prices[i]-fee)
	}
	return dp2[len(prices)-1]
}
