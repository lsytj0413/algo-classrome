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

package p0462

import "sort"

func minMoves2(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	sort.Ints(nums)
	mid := nums[len(nums)/2]
	sum := 0
	for _, v := range nums {
		if v > mid {
			sum = sum + v - mid
		} else {
			sum = sum + mid - v
		}
	}
	return sum
}
