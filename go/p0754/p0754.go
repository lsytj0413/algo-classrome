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

package p0754

func reachNumber(target int) int {
	if target < 0 {
		target = 0 - target
	}

	n, sum := 0, 0
	for sum < target {
		n++
		sum += n
	}

	dis := sum - target
	if dis&0x01 == 1 {
		if n&0x01 == 1 {
			n += 2
		} else {
			n++
		}
	}

	return n
}
