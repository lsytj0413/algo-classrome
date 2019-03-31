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

package p0503

import (
	"container/list"
	"math"
)

func nextGreaterElements(nums []int) []int {
	r := make([]int, len(nums))
	stack := list.New()

	for i := len(nums) - 1; i >= 0; i-- {
		for e := stack.Back(); e != nil && e.Value.(int) <= nums[i]; e = stack.Back() {
			stack.Remove(e)
		}

		if stack.Len() != 0 {
			e := stack.Back()
			r[i] = e.Value.(int)
		} else {
			r[i] = math.MinInt32
		}
		stack.PushBack(nums[i])
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if r[i] != math.MinInt32 {
			continue
		}

		for e := stack.Back(); e != nil && e.Value.(int) <= nums[i]; e = stack.Back() {
			stack.Remove(e)
		}

		if stack.Len() != 0 {
			e := stack.Back()
			r[i] = e.Value.(int)
		} else {
			r[i] = -1
		}
	}

	return r
}
