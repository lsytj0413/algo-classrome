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

package p0477

func totalHammingDistance(nums []int) int {
	hanming := [2]int{0, 0}
	r := 0
	for i := 0; i < 32; i++ {
		hanming[0], hanming[1] = 0, 0
		sum := 0
		for j, num := range nums {
			hanming[num&0x01]++
			sum += num
			nums[j] = num >> 1
		}

		if sum == 0 {
			break
		}

		r = r + hanming[0]*hanming[1]
	}
	return r
}
