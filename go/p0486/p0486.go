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

package p0486

// PredictTheWinner will predict the winner of nums
func PredictTheWinner(nums []int) bool {
	size := len(nums)
	dp := make([][]int, size)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, size)
		dp[i][i] = nums[i]
	}

	_max := func(v1 int, v2 int) int {
		if v1 > v2 {
			return v1
		}

		return v2
	}

	for i := 1; i < size; i++ {
		for j := 0; j < size-i; j++ {
			dp[j][j+i] = _max(nums[j+i]-dp[j][j+i-1], nums[j]-dp[j+1][j+i])
		}
	}

	return dp[0][size-1] >= 0
}
