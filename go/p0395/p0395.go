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

package p0395

import (
	"math"
)

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}

	ct := [26]int{}
	for i := 0; i < len(s); i++ {
		ct[int(s[i]-'a')]++
	}

	for i := 0; i < len(s); i++ {
		if ct[int(s[i]-'a')] < k {
			left := longestSubstring(s[0:i], k)
			right := longestSubstring(s[i+1:len(s)], k)
			return int(math.Max(float64(left), float64(right)))
		}
	}

	return len(s)
}
