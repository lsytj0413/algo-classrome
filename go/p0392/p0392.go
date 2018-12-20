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

package p0392

// TODO: 动态规划算法, 以及思考题中的预处理(t)算法
func isSubsequence(s string, t string) bool {
	// 1: 朴素解法
	_f1 := func(s string, t string) bool {
		j := 0
		for i := 0; i < len(t) && j < len(s); i++ {
			if s[j] == t[i] {
				j++
			}
		}

		return j == len(s)
	}
	_ = _f1

	return _f1(s, t)
}
