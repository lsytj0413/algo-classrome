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

package p0944

func minDeletionSize(A []string) int {
	res := 0
	for i := 0; i < len(A[0]); i++ {
		for j := 0; j < len(A)-1; j++ {
			if A[j][i] > A[j+1][i] {
				res++
				break
			}
		}
	}
	return res
}
