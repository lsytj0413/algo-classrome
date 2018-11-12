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

package p0347

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	maps := map[int]int{}
	for _, num := range nums {
		maps[num] = maps[num] + 1
	}

	type NumCount struct {
		num   int
		count int
	}
	numCounts := make([]NumCount, 0, len(maps))
	for num, count := range maps {
		numCounts = append(numCounts, NumCount{
			num:   num,
			count: count,
		})
	}

	// TODO: 最小堆可以进一步降低复杂度, 或者使用 quickSelect
	sort.Slice(numCounts, func(i int, j int) bool {
		return numCounts[i].count < numCounts[j].count
	})

	r := make([]int, 0, k)
	for j := len(numCounts) - 1; j >= 0 && k > 0; j, k = j-1, k-1 {
		r = append(r, numCounts[j].num)
	}

	return r
}
