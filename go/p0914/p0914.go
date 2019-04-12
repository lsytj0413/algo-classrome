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

package p0914

import "sort"

// TODO: optimize the arr
func hasGroupsSizeX(deck []int) bool {
	if len(deck) < 2 {
		return false
	}

	counts := make(map[int]int)
	for _, v := range deck {
		counts[v]++
	}

	arr := make([]int, 0, len(counts))
	for _, v := range counts {
		arr = append(arr, v)
	}
	sort.Ints(arr)

	min := arr[0]
	if min <= 1 {
		return false
	}
	arr = arr[1:]

	for i := 2; i <= min; i++ {
		canGroup := true
		for _, v := range arr {
			if v%i != 0 {
				canGroup = false
				break
			}
		}
		if canGroup {
			return true
		}
	}

	return false
}
