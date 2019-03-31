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

package p0526

// TODO: use A! algothrim
func countArrangement(N int) int {
	used, sum := make([]bool, N), 0

	var helper func(index int)
	helper = func(index int) {
		if index > N {
			sum++
			return
		}

		for i := 1; i <= len(used); i++ {
			if !used[i-1] && ((index%i == 0) || (i%index == 0)) {
				used[i-1] = true
				helper(index + 1)
				used[i-1] = false
			}
		}
	}
	helper(1)

	return sum
}
