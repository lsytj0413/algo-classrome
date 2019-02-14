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

package p0442

func findDuplicates(nums []int) []int {
	for i := 0; i < len(nums); {
		v := nums[i]
		switch {
		case v <= 0 || v == i+1:
			// the current number is in the right bulk
			i++
		default:
			if nums[v-1] == v {
				nums[v-1] = 0 - v
				nums[i] = 0
				i++
			} else {
				nums[i], nums[v-1] = nums[v-1], nums[i]
			}
		}
	}

	r := make([]int, 0)
	for _, v := range nums {
		if v < 0 {
			r = append(r, (0 - v))
		}
	}
	return r
}
