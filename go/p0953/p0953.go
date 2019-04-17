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

package p0953

func isAlienSorted(words []string, order string) bool {
	trans := make([]int, 26)
	for i := 0; i < len(order); i++ {
		trans[int(order[i]-'a')] = i
	}

	_less := func(x string, y string) bool {
		i := 0
		for i < len(x) && i < len(y) {
			ix, iy := trans[int(x[i]-'a')], trans[int(y[i]-'a')]
			switch {
			case ix > iy:
				return false
			case ix < iy:
				return true
			}
			i++
		}

		if len(x) > len(y) {
			return false
		}
		return true
	}

	n := len(words)
	for i := n - 1; i > 0; i-- {
		if _less(words[i], words[i-1]) {
			return false
		}
	}
	return true
}
