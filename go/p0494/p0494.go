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

package p0494

func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum < S || (S+sum)%2 != 0 {
		return 0
	}

	target := (S + sum) >> 1
	dp := make([]int, target+1)
	dp[0] = 1
	for _, v := range nums {
		for i := target; i > v-1; i-- {
			dp[i] += dp[i-v]
		}
	}
	return dp[target]
}
