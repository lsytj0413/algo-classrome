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

package p0970

import (
	"math"
	"sort"
)

// TODO: optimize
func powerfulIntegers(x int, y int, bound int) []int {
	marks := make(map[int]bool)
	for i := 0; ; i++ {
		vi := int(math.Pow(float64(x), float64(i)))
		if vi > bound || (i > 0 && x == 1) {
			break
		}
		for j := 0; ; j++ {
			v := vi + int(math.Pow(float64(y), float64(j)))
			if v > bound || (j > 0 && y == 1) {
				break
			}
			marks[v] = true
		}
	}

	r := make([]int, 0, len(marks))
	for k := range marks {
		r = append(r, k)
	}
	sort.Ints(r)
	return r
}
