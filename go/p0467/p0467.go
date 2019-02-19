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

package p0467

func findSubstringInWraproundString(p string) int {
	if len(p) == 0 {
		return 0
	}

	arr := make([]int, 26)
	pre := 1
	arr[int(p[0]-'a')] = 1

	for i := 1; i < len(p); i++ {
		c := 1
		if (p[i-1] == 'z' && p[i] == 'a') || (int(p[i]-p[i-1]) == 1) {
			c = pre + 1
		}
		pre = c

		if c > arr[int(p[i]-'a')] {
			arr[int(p[i]-'a')] = c
		}
	}

	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}
