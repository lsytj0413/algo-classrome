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

package p0456

import (
	"container/list"
	"math"
)

func find132pattern(nums []int) bool {
	stack := list.New()
	min := math.MinInt32

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < min {
			return true
		}

		for stack.Len() > 0 {
			e := stack.Back()
			v := e.Value.(int)
			if nums[i] > v {
				min = v
				stack.Remove(e)
			} else {
				break
			}
		}

		stack.PushBack(nums[i])
	}

	return false
}
