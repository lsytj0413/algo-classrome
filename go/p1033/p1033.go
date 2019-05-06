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

package p1033

import "sort"

func numMovesStones(a int, b int, c int) []int {
	arr := []int{a, b, c}
	sort.Ints(arr)

	if (arr[2] - arr[0]) == 2 {
		return []int{0, 0}
	}

	if ((arr[2] - arr[1]) <= 2) || ((arr[1] - arr[0]) <= 2) {
		return []int{1, arr[2] - arr[0] - 2}
	}

	return []int{2, arr[2] - arr[0] - 2}
}
