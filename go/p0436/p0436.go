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

package p0436

import (
	"sort"

	"github.com/lsytj0413/algo-classrome/go/comm"
)

// Interval is a struct with start and end point
type Interval = comm.Interval

// TODO: binary-search and rb-tree
func findRightInterval(intervals []Interval) []int {
	r := make([]int, 0, len(intervals))
	indexes := make(map[int]int, len(intervals))
	for i, v := range intervals {
		r = append(r, v.End)
		indexes[v.Start] = i
	}

	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	for i := 0; i < len(r); i++ {
		v := -1
		for j := 0; j < len(intervals); j++ {
			if intervals[j].Start >= r[i] {
				v = indexes[intervals[j].Start]
				break
			}
		}
		r[i] = v
	}

	return r
}
