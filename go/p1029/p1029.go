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

package p1029

import (
	"sort"
)

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i int, j int) bool {
		return (costs[i][0] - costs[i][1]) < (costs[j][0] - costs[j][1])
	})

	cost := 0
	for i := 0; i < len(costs)/2; i++ {
		cost += costs[i][0]
	}

	for i := len(costs) / 2; i < len(costs); i++ {
		cost += costs[i][1]
	}

	return cost
}
