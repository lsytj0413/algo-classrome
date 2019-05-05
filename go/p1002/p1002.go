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

package p1002

func commonChars(A []string) []string {
	if len(A) == 0 {
		return []string{}
	}

	_toArray := func(str string) []int {
		res := make([]int, 26)
		for i := 0; i < len(str); i++ {
			res[int(str[i]-'a')]++
		}

		return res
	}

	_min := func(x int, y int) int {
		if x < y {
			return x
		}
		return y
	}

	res := _toArray(A[0])
	for i := 1; i < len(A); i++ {
		arr := _toArray(A[i])

		for j := 0; j < len(arr); j++ {
			if res[j] == 0 || arr[j] == 0 {
				res[j] = 0
			} else {
				res[j] = _min(res[j], arr[j])
			}
		}
	}

	ret := make([]string, 0)
	for i := 0; i < len(res); i++ {
		for res[i] > 0 {
			ret = append(ret, string([]byte{byte(i) + 'a'}))
			res[i]--
		}
	}
	return ret
}
