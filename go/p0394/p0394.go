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

package p0394

import (
	"strings"
)

func decodeString(s string) string {
	r, num, depth, sub := strings.Builder{}, 0, 0, strings.Builder{}

	for i := 0; i < len(s); i++ {
		if depth > 0 && s[i] != ']' {
			sub.WriteByte(s[i])
			if s[i] == '[' {
				depth++
			}
			continue
		}

		switch {
		case s[i] >= '0' && s[i] <= '9':
			num = num*10 + int(s[i]-'0')
		case s[i] == '[':
			depth++
		case s[i] == ']':
			depth--
			if depth == 0 {
				v := decodeString(sub.String())
				for j := 0; j < num; j++ {
					r.WriteString(v)
				}
				num = 0
				sub.Reset()
			} else {
				sub.WriteByte(s[i])
			}
		default:
			r.WriteByte(s[i])
		}
	}

	return r.String()
}
