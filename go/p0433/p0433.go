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

package p0433

import (
	"container/list"
)

func minMutation(start string, end string, bank []string) int {
	digits := []byte{'A', 'C', 'G', 'T'}
	bankSet := make(map[string]bool, len(bank))
	for _, v := range bank {
		bankSet[v] = true
	}

	if !bankSet[end] {
		return -1
	}
	if bankSet[start] && start == end {
		return 0
	}

	visited := make(map[string]bool, len(bank))
	visited[start] = true
	l := list.New()
	l.PushBack(start)
	min := 0

	for l.Len() > 0 {
		n := l.Len()
		for n > 0 {
			e := l.Front()
			v := e.Value.(string)
			for i := 0; i < len(v); i++ {
				bytev := []byte(v)
				for _, digit := range digits {
					if digit == bytev[i] {
						continue
					}

					bytev[i] = digit
					r := string(bytev)
					if r == end {
						return min + 1
					}

					if bankSet[r] && !visited[r] {
						l.PushBack(r)
						visited[r] = true
					}
				}
			}
			l.Remove(e)
			n--
		}
		min++
	}

	return -1
}
