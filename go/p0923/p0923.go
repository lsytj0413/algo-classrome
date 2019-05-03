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

package p0923

func threeSumMulti(A []int, target int) int {
	counts := make([]int, 101)
	for _, v := range A {
		counts[v]++
	}

	res := 0
	for i := 0; i <= target; i++ {
		for j := i; j <= target; j++ {
			k := target - i - j
			if k < 0 || k > 100 || k < j {
				continue
			}
			if counts[k] == 0 || counts[i] == 0 || counts[j] == 0 {
				continue
			}

			switch {
			case i == j && j == k:
				res = res + (counts[i]*(counts[i]-1)*(counts[i]-2))/6
			case i == j && j != k:
				res = res + (counts[i]*(counts[i]-1)*counts[k])/2
			case i != j && j == k:
				res = res + (counts[j]*(counts[j]-1)*counts[i])/2
			default:
				res = res + (counts[i] * counts[j] * counts[k])
			}
		}
	}
	return res % (1000000007)
}
