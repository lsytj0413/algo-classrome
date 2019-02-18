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

package p0424

func characterReplacement(s string, k int) int {
	if len(s) <= k {
		return len(s)
	}

	bits := make([]int, 26)

	_most := func(arr []int) int {
		max := 0
		for _, v := range arr {
			if v > max {
				max = v
			}
		}
		return max
	}

	_len := func(start int, end int) int {
		v := end - start + 1
		return v
	}

	start, end, max := 0, 0, k+1
	for end <= k {
		bits[int(s[end]-'A')]++
		end++
	}

	for end < len(s) {
		l := _len(start, end)
		bits[int(s[end]-'A')]++
		n := _most(bits)
		switch {
		case l <= k || (l-n) <= k:
			if l > max {
				max = l
			}
			end++
		default:
			bits[int(s[start]-'A')]--
			bits[int(s[end]-'A')]--
			start++
		}
	}
	return max
}
