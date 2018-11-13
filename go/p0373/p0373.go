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

package p0373

import (
	"math"
)

// TODO: 最大堆解法
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	cnt := k
	if len(nums1)*len(nums2) < cnt {
		cnt = len(nums1) * len(nums2)
	}
	r := make([][]int, 0, cnt)
	index := make([]int, len(nums1))
	for cnt > 0 {
		cnt--
		min, m := math.MaxInt32, 0
		for i := 0; i < len(nums1); i++ {
			if index[i] < len(nums2) && nums1[i]+nums2[index[i]] <= min {
				min = nums1[i] + nums2[index[i]]
				m = i
			}
		}
		r = append(r, []int{nums1[m], nums2[index[m]]})
		index[m] = index[m] + 1
	}

	return r
}
