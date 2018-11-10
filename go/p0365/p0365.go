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

package p0365

func canMeasureWater(x int, y int, z int) bool {
	if z > x+y {
		return false
	}
	if z == 0 {
		return true
	}

	var _gcd func(int, int) int
	_gcd = func(x int, y int) int {
		if y == 0 {
			return x
		}
		return _gcd(y, x%y)
	}

	return z%_gcd(x, y) == 0
}
