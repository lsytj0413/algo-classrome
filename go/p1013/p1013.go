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

package p1013

func canThreePartsEqualSum(A []int) bool {
	if len(A) < 3 {
		return false
	}

	sum := 0
	for _, v := range A {
		sum += v
	}
	if sum%3 != 0 {
		return false
	}
	sum = sum / 3

	i, sumi := 1, A[0]
	for sumi != sum && i < len(A)-1 {
		sumi += A[i]
		i++
	}
	if i >= len(A)-1 {
		return false
	}

	j, sumj := len(A)-2, A[len(A)-1]
	for sumj != sum && j > 0 {
		sumj += A[j]
		j--
	}

	return i <= j
}
