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

package p0397

func integerReplacement(n int) int {
	_min := func(x int, y int) int {
		if x > y {
			return y
		}
		return x
	}

	_f1 := func(n int) int {
		dp := make([]int, n+2)
		dp[1] = 0
		i := 2
		for i <= n {
			switch i & 1 {
			case 0: // even number
				dp[i] = dp[i>>1] + 1
			default: // odd number
				dp[i+1] = dp[(i+1)>>1] + 1
				dp[i] = _min(dp[i-1], dp[i+1]) + 1
			}
			i++
		}
		return dp[n]
	}
	_ = _f1

	_f2 := func(n int) int {
		count := 0
		for n > 1 {
			if n&1 == 0 {
				n = n >> 1
			} else {
				// if (n+1)%4 == 0 && (n-1) != 2 {
				if n&0x03 == 3 && n != 3 {
					n++
				} else {
					n--
				}
			}
			count++
		}
		return count
	}
	_ = _f2

	return _f2(n)
}
