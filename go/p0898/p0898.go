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

package p0898

func subarrayBitwiseORs(A []int) int {
	xor := make([]int, len(A))
	xor[0] = A[0]
	sets := map[int]bool{}
	for _, v := range A {
		sets[v] = true
	}

	for i := 1; i < len(A); i++ {
		xor[i] = A[i]
		sets[xor[i]] = true
		for j := i - 1; j >= 0; j-- {
			v := A[i] | xor[j]
			if v == xor[j] {
				break
			}

			xor[j] = v
			sets[v] = true
		}
	}

	return len(sets)
}
