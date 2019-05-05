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

package p1005

import (
	"sort"
)

// TODO: optimize the code
func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)

	i := 0
	for K > 0 {
		if i < len(A) && A[i] < 0 {
			A[i] = -A[i]
			i++
			K--
			continue
		}

		switch {
		case K&0x01 == 0:
		case i == len(A):
			A[len(A)-1] = -A[len(A)-1]
		case i == 0:
			A[0] = -A[0]
		default:
			if A[i] > A[i-1] {
				A[i-1] = -A[i-1]
			} else {
				A[i] = -A[i]
			}
		}
		break
	}

	res := 0
	for _, v := range A {
		res += v
	}
	return res
}
