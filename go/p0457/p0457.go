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

package p0457

// TODO: 应该可以不使用快慢指针方式? 使用 nums 自身来记录是否访问过节点即可
func circularArrayLoop(nums []int) bool {
	next := func(nums []int, i int) int {
		length := len(nums)

		if length == 1 {
			return 0
		}

		np := i + nums[i]
		if np >= 0 {
			return np % length
		}

		return length + np%length
	}

	for i := 0; i < len(nums); i++ {
		slow, fast := i, next(nums, i)
		for nums[i]*nums[slow] > 0 && nums[i]*nums[fast] > 0 && nums[i]*nums[next(nums, fast)] > 0 && nums[slow] != 0 && nums[fast] != 0 {
			if slow == fast {
				if slow == next(nums, slow) {
					break
				}
				return true
			}

			slow, fast = next(nums, slow), next(nums, next(nums, fast))
		}

		nums[i] = 0
	}

	return false
}
