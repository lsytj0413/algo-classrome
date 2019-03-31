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

package p0473

import (
	"sort"
)

func makesquare(nums []int) bool {
	if len(nums) < 4 {
		return false
	}

	sum, max := 0, 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}

	if sum%4 != 0 || max > (sum/4) {
		return false
	}

	var findOne func(int, []int) bool
	findOne = func(target int, nums []int) bool {
		if target == 0 {
			return true
		}
		for i := 0; i < len(nums); i++ {
			if nums[i] <= 0 {
				continue
			}

			if target < nums[i] {
				continue
			}

			n := nums[i]
			nums[i] = 0
			if target == n {
				return true
			}

			ok := findOne(target-n, nums)
			if ok {
				return true
			}
			nums[i] = n
		}
		return false
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	for i := 0; i < 4; i++ {
		ok := findOne(sum/4, nums)
		if !ok {
			return false
		}
	}

	return true
}
