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

package p0452

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i int, j int) bool {
		if points[i][0] < points[j][0] {
			return true
		}
		return false
	})

	count, lastEnd := 1, points[0][1]
	for i := 0; i < len(points); i++ {
		if lastEnd < points[i][0] {
			count++
			lastEnd = points[i][1]
			continue
		}

		if lastEnd > points[i][1] {
			lastEnd = points[i][1]
		}
	}

	return count
}
