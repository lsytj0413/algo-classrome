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

package p0491

import (
	"fmt"
	"strings"
)

func findSubsequences(nums []int) [][]int {
	set := make(map[string]bool)

	_toString := func(arr []int) string {
		b := strings.Builder{}
		for _, v := range arr {
			b.WriteString(fmt.Sprintf("%v", v))
			b.WriteString(".")
		}

		return b.String()
	}

	r := make([][]int, 0)
	_addToR := func(arr []int) {
		v := _toString(arr)
		if !set[v] {
			r = append(r, arr)
			set[v] = true
		}
	}

	for i := 1; i < len(nums); i++ {
		rl := len(r)
		for j := 0; j < rl; j++ {
			if nums[i] >= r[j][len(r[j])-1] {
				cp := make([]int, len(r[j]))
				copy(cp, r[j])
				_addToR(append(cp, nums[i]))
			}
		}

		for j := 0; j < i; j++ {
			if nums[i] >= nums[j] {
				_addToR([]int{nums[j], nums[i]})
			}
		}
	}

	return r
}
