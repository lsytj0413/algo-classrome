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

package p0985

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	prev := 0
	for _, v := range A {
		if v&0x01 == 0 {
			prev += v
		}
	}

	res := make([]int, 0, len(queries))
	for _, query := range queries {
		switch {
		case query[0]&0x01 == 0 && A[query[1]]&0x01 == 0:
			prev += query[0]
		case query[0]&0x01 == 1 && A[query[1]]&0x01 == 0:
			prev -= A[query[1]]
		case query[0]&0x01 == 0 && A[query[1]]&0x01 == 1:
			// nothing
		case query[0]&0x01 == 1 && A[query[1]]&0x01 == 1:
			prev = prev + query[0] + A[query[1]]
		}

		A[query[1]] += query[0]
		res = append(res, prev)
	}
	return res
}
