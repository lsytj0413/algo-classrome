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

package p1014

func maxScoreSightseeingPair(A []int) int {
	if len(A) < 2 {
		return 0
	}

	max, res := A[0], 0
	for i := 1; i < len(A); i++ {
		res0 := A[i] - i + max
		if res0 > res {
			res = res0
		}

		v := A[i] + i
		if v > max {
			max = v
		}
	}

	return res
}
