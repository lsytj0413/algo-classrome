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

package p0388

import (
	"strings"
)

func lengthLongestPath(input string) int {
	input = strings.Replace(input, "    ", "\t", -1) + "\n"
	max, last, nt, file := 0, -1, 0, false
	depth := make([]int, 0)
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '\n':
			name := string(input[last+1 : i])

			if nt < len(depth) {
				depth = depth[0:nt]
			}

			if file {
				d := len(name) + len(depth)
				for j := 0; j < len(depth); j++ {
					d += depth[j]
				}

				if nt > len(depth) {
					d += (nt - len(depth)) * 4
				}

				if d > max {
					max = d
				}
			} else {
				depth = append(depth, len(name))
			}
			file = false
			nt = 0
			last = i

		case '\t':
			nt++
			last = i
		case '.':
			file = true
		}
	}

	return max
}
