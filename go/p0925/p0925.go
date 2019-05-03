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

package p0925

// TODO: optimize the code
func isLongPressedName(name string, typed string) bool {
	if len(typed) == 0 {
		return len(name) == 0
	}
	if len(name) == 0 {
		return false
	}

	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		if name[i] == typed[j] {
			i++
			j++
			continue
		}

		if j == 0 || typed[j] != typed[j-1] {
			return false
		}
		j++
	}

	if i != len(name) {
		return false
	}

	for j < len(typed) {
		if j > 0 {
			if typed[j] != typed[j-1] {
				return false
			}
		}

		j++
	}
	return true
}
