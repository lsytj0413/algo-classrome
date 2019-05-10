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

package p0989

// TODO: optimize the slice append
func addToArrayForm(A []int, K int) []int {
	for i := len(A) - 1; i >= 0 && (K) != 0; i-- {
		v := A[i] + K%10
		K = K/10 + v/10
		A[i] = v % 10
	}

	if K == 0 {
		return A
	}

	res := make([]int, 0)
	for K > 0 {
		res = append(res, K%10)
		K = K / 10
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	res = append(res, A...)
	return res
}
