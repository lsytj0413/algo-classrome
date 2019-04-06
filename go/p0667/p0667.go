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

package p0667

func constructArray(n int, k int) []int {
	r := make([]int, 0, n)
	l, h := 1, n
	for l <= h {
		if k > 1 {
			if k&0x01 == 1 {
				r = append(r, l)
				l++
			} else {
				r = append(r, h)
				h--
			}
			k--
			continue
		}

		r = append(r, l)
		l++
	}
	return r
}
