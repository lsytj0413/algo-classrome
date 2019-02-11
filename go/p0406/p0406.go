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

package p0406

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i int, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}

		return people[i][0] > people[j][0]
	})

	r := make([][]int, len(people))
	for i := 0; i < len(people); i++ {
		index := people[i][1]
		for j := len(people) - 1; j >= index && j > 0; j-- {
			r[j] = r[j-1]
		}

		r[index] = people[i]
	}
	return r
}
