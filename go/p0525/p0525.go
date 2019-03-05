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

package p0525

func findMaxLength(nums []int) int {
	for i := range nums {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}

	m := make(map[int]int)
	sum, ans := 0, 0
	for i := range nums {
		sum += nums[i]
		if sum == 0 {
			if i >= ans {
				ans = i + 1
			}
		}

		if _, ok := m[sum]; !ok {
			m[sum] = i
			continue
		}

		v := i - m[sum]
		if v > ans {
			ans = v
		}
	}

	return ans
}
