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

package p0948

import (
	"sort"
)

func bagOfTokensScore(tokens []int, P int) int {
	sort.Ints(tokens)
	res, now := 0, 0
	i, j := 0, len(tokens)-1

	for i <= j {
		if P >= tokens[i] {
			now++
			if now > res {
				res = now
			}
			P -= tokens[i]
			i++
			continue
		}
		if now <= 0 {
			break
		}

		now--
		P += tokens[j]
		j--
	}

	return res
}
