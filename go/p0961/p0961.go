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

package p0961

func repeatedNTimes(A []int) int {
	for i := 0; i < len(A)-1; i++ {
		if A[i] == A[i+1] {
			return A[i]
		}
	}

	for i := 2; i < len(A); i++ {
		if A[i] == A[0] {
			return A[0]
		}
		if A[i] == A[1] {
			return A[1]
		}
	}

	return A[0]
}
