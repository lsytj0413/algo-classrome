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
			} else if weights[i] <= ship {
				now = weights[i]
				v++
			} else {
				return false
			}

			if v > D {
				return false
			}
		}

		return true
	}

	min, max, res := res, sum, 0
	for min < max {
		res = (min + max) / 2
		if _can(weights, res, D) {
			max = res - 1
		} else {
			min = res + 1
		}
	}

	if _can(weights, max, D) {
		return max
	}
	return max + 1
}
