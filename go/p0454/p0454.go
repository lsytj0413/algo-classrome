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

package p0454

func fourSumCount(A []int, B []int, C []int, D []int) int {
	counts := make(map[int]int)
	for _, va := range A {
		for _, vb := range B {
			counts[va+vb] = counts[va+vb] + 1
		}
	}

	r := 0
	for _, vc := range C {
		for _, vd := range D {
			v := 0 - vc - vd
			r += counts[v]
		}
	}

	return r
}
