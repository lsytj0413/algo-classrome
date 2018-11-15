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

package p0376

// TODO: 动态规划
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	r, flag := len(nums), 0
	for i := 1; i < len(nums); i++ {
		switch {
		case nums[i]-nums[i-1] == 0:
			r--
		case nums[i]-nums[i-1] > 0:
			if flag == 1 {
				r--
			} else {
				flag = 1
			}
		default:
			if flag == -1 {
				r--
			} else {
				flag = -1
			}
		}
	}

	return r
}
