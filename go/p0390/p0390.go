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

package p0390

func lastRemaining(n int) int {
	var _helper func(int, bool) int
	_helper = func(n int, isLeft bool) int {
		if n == 1 {
			return 1
		}

		if isLeft {
			return 2 * _helper(n/2, false)
		}
		return 2*_helper(n/2, true) - 1 + n%2
	}

	return _helper(n, true)
}
