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

package p0332

import (
	"container/list"
	"sort"
)

func findItinerary(tickets [][]string) []string {
	maps := make(map[string][]string)
	for _, ticket := range tickets {
		maps[ticket[0]] = append(maps[ticket[0]], ticket[1])
	}
	for k, v := range maps {
		sort.Strings(v)
		maps[k] = v
	}

	roads := list.New()
	roads.PushBack("JFK")
	r := []string{}

	for roads.Len() != 0 {
		e := roads.Back()
		v := e.Value.(string)
		targets := maps[v]
		if targets == nil || len(targets) == 0 {
			roads.Remove(e)
			r = append(r, v)
		} else {
			roads.PushBack(targets[0])
			maps[v] = targets[1:]
		}
	}

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}
