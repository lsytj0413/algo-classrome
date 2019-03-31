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

package p0481

func magicalString(n int) int {
	switch {
	case n == 0:
		return 0
	case n >= 1 && n <= 3:
		return 1
	}

	count, num, head, tail := 1, 1, 2, 3
	arr := make([]int, n+1)
	arr[0], arr[1], arr[2] = 1, 2, 2

	for tail < n {
		for i := 0; i < arr[head]; i++ {
			arr[tail] = num
			if tail < n && num == 1 {
				count++
			}
			tail++
		}

		num = num ^ 3
		head++
	}

	return count
}
