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

package p0377

// TODO: 另一种迭代 nums 的方法
func combinationSum4(nums []int, target int) int {
	maps := make(map[int]bool, len(nums))
	for _, num := range nums {
		maps[num] = true
	}

	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		sum := 0
		for j := 0; j <= i; j++ {
			if maps[i-j] {
				sum += dp[j]
			}
		}
		dp[i] = sum
	}

	return dp[target]
}
