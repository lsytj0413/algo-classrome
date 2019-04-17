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

package p0977

func sortedSquares(A []int) []int {
	r := make([]int, len(A))
	w := len(r) - 1
	i, j := 0, len(A)-1

	_abs := func(x int) int {
		if x < 0 {
			return 0 - x
		}
		return x
	}

	for w >= 0 {
		vi, vj := _abs(A[i]), _abs(A[j])
		if vi < vj {
			r[w] = vj * vj
			j--
		} else {
			r[w] = vi * vi
			i++
		}
		w--
	}

	return r
}
