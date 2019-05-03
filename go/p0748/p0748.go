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

package p0748

import (
	"strings"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	counts := func(v string) []int {
		c := make([]int, 26)
		v = strings.ToLower(v)
		for i := 0; i < len(v); i++ {
			if v[i] >= 'a' && v[i] <= 'z' {
				c[int(v[i]-'a')]++
			}
		}
		return c
	}
	licenseCounts := counts(licensePlate)

	_isComplete := func(v []int, license []int) bool {
		for i := 0; i < len(v); i++ {
			if v[i] < license[i] && license[i] > 0 {
				return false
			}
		}
		return true
	}
	res := ""
	for _, word := range words {
		if res == "" || len(res) > len(word) {
			v := counts(word)
			if _isComplete(v, licenseCounts) {
				res = word
			}
		}
	}

	return res
}
