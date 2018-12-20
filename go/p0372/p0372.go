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

package p0372

// TODO: 耗时更短的方式
func superPow(a int, b []int) int {
	r := 1
	a = a % 1337

	_isZero := func(b []int) bool {
		for i := 0; i < len(b); i++ {
			if b[i] > 0 {
				return false
			}
		}

		return true
	}

	_div := func(b []int) {
		tmp := 0
		for i := 0; i < len(b); i++ {
			b[i] += tmp * 10
			tmp = b[i] % 2
			b[i] = b[i] / 2
		}
	}

	for !_isZero(b) {
		if b[len(b)-1]%2 != 0 {
			r = (r * a) % 1337
		}

		_div(b)
		a = (a * a) % 1337
	}

	return r
}
