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

package p0942

func diStringMatch(S string) []int {
	min, max := 0, len(S)
	r := make([]int, len(S)+1)

	for i := 0; i < len(S); i++ {
		switch S[i] {
		case 'I':
			r[i] = min
			min++
		case 'D':
			r[i] = max
			max--
		}
	}
	r[len(r)-1] = min

	return r
}
