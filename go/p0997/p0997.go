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

package p0997

// TODO: 图算法?
func findJudge(N int, trust [][]int) int {
	marks, people := make(map[int]bool, N), (N*(N+1))>>1
	for _, v := range trust {
		if !marks[v[0]] {
			people -= v[0]
			marks[v[0]] = true
		}
	}

	marks = make(map[int]bool, N)
	for _, pair := range trust {
		if pair[1] == people {
			marks[pair[0]] = true
		}
	}

	if len(marks) == N-1 {
		return people
	}

	return -1
}
