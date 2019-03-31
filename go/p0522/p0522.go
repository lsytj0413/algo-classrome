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

package p0522

import (
	"sort"
)

func findLUSlength(strs []string) int {
	sort.Slice(strs, func(i int, j int) bool {
		return len(strs[i]) > len(strs[j])
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

	for i := 0; i < len(strs); i++ {
		found := false
		for j := 0; j < len(strs); j++ {
			if i == j {
				continue
			}
			if len(strs[i]) > len(strs[j]) {
				break
			}
			if !_isUnique(strs[j], strs[i]) {
				found = true
				break
			}
		}

		if !found {
			return len(strs[i])
		}
	}

	return -1
}
