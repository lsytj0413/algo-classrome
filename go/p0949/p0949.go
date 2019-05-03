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

package p0949

// TODO: optimize the code
func largestTimeFromDigits(A []int) string {
	used := []bool{false, false, false, false}
	r := make([]byte, 5)

	_has := func(v int) bool {
		for i := 0; i < len(A); i++ {
			if A[i] == v && used[i] == false {
				return true
			}
		}
		return false
	}

	_use := func(v int) {
		for i := 0; i < len(A); i++ {
			if A[i] == v && used[i] == false {
				used[i] = true
				return
			}
		}
	}
	_unuse := func(v int) {
		for i := 0; i < len(A); i++ {
			if A[i] == v && used[i] == true {
				used[i] = false
				return
			}
		}
	}

	_minute := func() bool {
		for i := 5; i >= 0; i-- {
			if _has(i) {
				_use(i)
				r[3] = '0' + byte(i)
				for j := 9; j >= 0; j-- {
					if _has(j) {
						r[4] = '0' + byte(j)
						return true
					}
				}
				_unuse(i)
			}
		}
		return false
	}

	r[2] = ':'
	if _has(2) {
		_use(2)
		r[0] = '2'
		for i := 3; i >= 0; i-- {
			if _has(i) {
				_use(i)
				r[1] = '0' + byte(i)
				if _minute() {
					return string(r)
				}
				_unuse(i)
			}
		}
	}

	used = []bool{false, false, false, false}
	r = make([]byte, 5)
	r[2] = ':'

	if _has(1) {
		_use(1)
		r[0] = '1'
		for i := 9; i >= 0; i-- {
			if _has(i) {
				_use(i)
				r[1] = '0' + byte(i)
				if _minute() {
					return string(r)
				}
				_unuse(i)
			}
		}
	}

	used = []bool{false, false, false, false}
	r = make([]byte, 5)
	r[2] = ':'

	if _has(0) {
		_use(0)
		r[0] = '0'
		for i := 9; i >= 0; i-- {
			if _has(i) {
				_use(i)
				r[1] = '0' + byte(i)
				if _minute() {
					return string(r)
				}
				_unuse(i)
			}
		}
	}

	return ""
}
