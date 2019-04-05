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

package p0917

// TODO: use reflect to avoid cp
func reverseOnlyLetters(S string) string {
	v := []byte(S)

	_isAlpha := func(b byte) bool {
		return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
	}

	i, j := 0, len(v)-1
	for {
		for i < len(v) && !_isAlpha(v[i]) {
			i++
		}
		for j >= 0 && !_isAlpha(v[j]) {
			j--
		}
		if i >= j {
			break
		}

		v[i], v[j] = v[j], v[i]
		i++
		j--
	}

	return string(v)
}
