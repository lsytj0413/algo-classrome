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

package p0451

import (
	"bytes"
)

func frequencySort(s string) string {
	fs := make([]int, 256)
	for i := 0; i < len(s); i++ {
		fs[int(s[i])]++
	}

	max := 0
	for _, v := range fs {
		if v > max {
			max = v
		}
	}

	chars := make([][]byte, max+1)
	for i, v := range fs {
		if v > 0 {
			chars[v] = append(chars[v], byte(i))
		}
	}

	r := bytes.Buffer{}
	for i := len(chars) - 1; i > 0; i-- {
		if chars[i] == nil {
			continue
		}

		for _, c := range chars[i] {
			count := i
			for count > 0 {
				r.WriteByte(c)
				count--
			}
		}
	}
	return r.String()
}
