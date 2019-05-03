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

package p0672

func flipLights(n int, m int) int {
	switch {
	case m == 0:
		return 1
	case n == 1:
		return 2
	case n == 2:
		if m == 1 {
			return 3
		}
		return 4
	case m == 1:
		return 4
	case m == 2:
		return 7
	default:
		return 8
	}
}
