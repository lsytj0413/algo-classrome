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

package p0929

// TODO: optimize
func numUniqueEmails(emails []string) int {
	normalize := func(v string) string {
		b := make([]byte, 0, len(v))
		skip, domain := false, false
		for i := 0; i < len(v); i++ {
			if domain {
				b = append(b, v[i])
			} else {
				switch v[i] {
				case '.':
				case '+':
					skip = true
				case '@':
					domain = true
					b = append(b, v[i])
				default:
					if !skip {
						b = append(b, v[i])
					}
				}
			}
		}
		return string(b)
	}

	counts := make(map[string]bool)
	for _, email := range emails {
		counts[normalize(email)] = true
	}
	return len(counts)
}
