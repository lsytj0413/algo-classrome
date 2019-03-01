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

package p0421

func findMaximumXOR(nums []int) int {
	mask, r := 0, 0
	for bit := 31; bit >= 0; bit-- {
		mask = mask | (1 << uint(bit))
		prefix := make(map[int]bool, len(nums)/(32-bit))
		for _, v := range nums {
			prefix[v&mask] = true
		}
		guess := r | (1 << uint(bit))
		for k := range prefix {
			if _, ok := prefix[k^guess]; ok {
				r = guess
				break
			}
		}
	}
	return r
}
