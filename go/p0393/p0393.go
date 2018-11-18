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

package p0393

func validUtf8(data []int) bool {
	next := 0
	for _, n := range data {
		if next > 0 {
			if (uint8(n)>>6)&0x03 != 2 {
				return false
			}
			next--
		} else {
			switch v := (uint8(n) >> 3) & 0x1F; {
			case v == 30:
				next = 3
			case v == 28 || v == 29:
				next = 2
			case v == 24 || v == 25 || v == 26 || v == 27:
				next = 1
			case v < 16:
				next = 0
			default:
				return false
			}
		}
	}

	return next == 0
}
