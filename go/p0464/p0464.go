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

package p0464

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger >= desiredTotal {
		return true
	}
	if (maxChoosableInteger*(maxChoosableInteger+1))/2 < desiredTotal {
		return false
	}

	record := make(map[int]bool)

	var canWin func(target int, visited int) bool
	canWin = func(target int, visited int) bool {
		if v, ok := record[visited]; ok {
			return v
		}

		for i := 1; i <= maxChoosableInteger; i++ {
			mask := (1 << uint(i))
			if ((mask & visited) == 0) && ((i >= target) || canWin(target-i, mask|visited) == false) {
				record[visited] = true
				return true
			}
		}

		record[visited] = false
		return false
	}

	return canWin(desiredTotal, 0)
}
