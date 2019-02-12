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

package p0423

// TODO: 用代码展开(现在是手动展开)
func originalDigits(s string) string {
	digitMaps := make(map[byte]int, 26)
	for i := 0; i < len(s); i++ {
		digitMaps[byte(s[i])] = digitMaps[byte(s[i])] + 1
	}

	_remove := func(v string) {
		for i := 0; i < len(v); i++ {
			digitMaps[byte(v[i])] = digitMaps[byte(v[i])] - 1
		}
	}

	digits := [10]int{}

	// single
	for digitMaps['z'] > 0 {
		_remove("zero")
		digits[0]++
	}

	for digitMaps['w'] > 0 {
		_remove("two")
		digits[2]++
	}

	for digitMaps['u'] > 0 {
		_remove("four")
		digits[4]++
	}

	for digitMaps['g'] > 0 {
		_remove("eight")
		digits[8]++
	}

	for digitMaps['x'] > 0 {
		_remove("six")
		digits[6]++
	}

	// multi
	for digitMaps['f'] > 0 {
		_remove("five")
		digits[5]++
	}

	for digitMaps['r'] > 0 {
		_remove("three")
		digits[3]++
	}

	for digitMaps['o'] > 0 {
		_remove("one")
		digits[1]++
	}

	for digitMaps['v'] > 0 {
		_remove("seven")
		digits[7]++
	}

	for digitMaps['i'] > 0 {
		_remove("nine")
		digits[9]++
	}

	r := ""
	for i := 0; i < 10; i++ {
		for digits[i] > 0 {
			r += string('0' + byte(i))
			digits[i]--
		}
	}
	return r
}
