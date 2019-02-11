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

package p0416

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum&0x01 == 1 {
		return false
	}

	C := sum/2 + 1
	dp := make([]bool, C)
	dp[0] = true

	for _, v := range nums {
		if C-1-v >= 0 && dp[C-1-v] {
			return true
		}

		for i := C - 2; i >= 0; i-- {
			if i-v >= 0 && dp[i-v] {
				dp[i] = true
			}
		}
	}

	return false
}
