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

package p0735

import (
	"container/list"
)

// TODO: avoid the list
func asteroidCollision(asteroids []int) []int {
	if len(asteroids) <= 1 {
		return asteroids
	}

	_abs := func(x int) int {
		if x < 0 {
			return 0 - x
		}
		return x
	}

	l := list.New()
	l.PushBack(asteroids[0])

	for i := 1; i < len(asteroids); i++ {
		store := false
		for {
			if l.Len() <= 0 {
				store = true
				break
			}

			e := l.Back()
			v := e.Value.(int)
			if !(v > 0 && asteroids[i] < 0) {
				store = true
				break
			}

			diff := _abs(asteroids[i]) - _abs(v)
			if diff == 0 {
				l.Remove(e)
				store = false
				break
			}
			if diff < 0 {
				store = false
				break
			}

			l.Remove(e)
			store = true
		}

		if store {
			l.PushBack(asteroids[i])
		}
	}

	r := make([]int, 0, l.Len())
	for l.Len() > 0 {
		e := l.Front()
		l.Remove(e)
		r = append(r, e.Value.(int))
	}
	return r
}
