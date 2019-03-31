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

package p0524

import (
	"sort"
)

// TODO: avoid sort twice/sort
func findLongestWord(s string, d []string) string {
	sort.Strings(d)
	sort.SliceStable(d, func(i int, j int) bool {
		if len(d[i]) == len(d[j]) {
			return false
		}

		return len(d[i]) > len(d[j])
	})

	_isUnique := func(s1 string, s2 string) bool {
		i, j := 0, 0
		for i < len(s1) && j < len(s2) {
			if s1[i] == s2[j] {
				i++
				j++
			} else {
				i++
			}
		}
		return j != len(s2)
	}

	for _, v := range d {
		if !_isUnique(s, v) {
			return v
		}
	}

	return ""
}
