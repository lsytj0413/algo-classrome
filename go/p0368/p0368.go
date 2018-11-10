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

package p0368

import (
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)

	dp, indexes, maxLength, index := make([]int, len(nums)), make([]int, len(nums)), 0, -1
	for i := range indexes {
		indexes[i] = -1
	}

	for i := range nums {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					indexes[i] = j
				}
			}
		}

		if dp[i] > maxLength {
			maxLength = dp[i]
			index = i
		}
	}

	r := make([]int, 0, maxLength)
	for index != -1 {
		r = append(r, nums[index])
		index = indexes[index]
	}

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}
