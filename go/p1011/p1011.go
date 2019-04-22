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

package p1011

import (
	"math"
)

// TODO: optimize code, and binary search
func shipWithinDays(weights []int, D int) int {
	sum, res := 0, 0
	for _, v := range weights {
		sum += v
		if v > res {
			res = v
		}
	}

	_can := func(weights []int, ship int, D int) bool {
		v, now := 0, math.MaxInt32
		for i := 0; i < len(weights); i++ {
			if now+weights[i] <= ship {
				now += weights[i]
			} else {
				now = weights[i]
				v++
			}
		}

		return v <= D
	}

	for res < sum {
		if _can(weights, res, D) {
			return res
		}
		res++
	}

	return sum
}
