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

package p0955

import (
	"container/list"
)

// TODO: optimize the code
func minDeletionSize(A []string) int {
	type Range struct {
		begin int
		end   int
	}
	l1, l2 := list.New(), list.New()
	l1.PushBack(Range{
		begin: 0,
		end:   len(A) - 1,
	})

	res := 0
	for i := 0; i < len(A[0]); i++ {
		mode := 0
		for e := l1.Front(); e != nil; e = e.Next() {
			v := e.Value.(Range)
			start := -1
			for j := v.begin; j < v.end; j++ {
				if A[j][i] > A[j+1][i] {
					mode = 2
					break
				} else if A[j][i] == A[j+1][i] {
					mode = 1
					if start == -1 {
						start = j
					}
					if j == v.end-1 {
						l2.PushBack(Range{
							begin: start,
							end:   v.end,
						})
					}
				} else {
					if start != -1 {
						l2.PushBack(Range{
							begin: start,
							end:   j,
						})
						start = -1
					}
				}
			}
			if mode == 2 {
				break
			}
		}

		switch mode {
		case 2:
			res++
		case 1:
			l1, l2 = l2, l1
		case 0:
			return res
		}
		l2 = list.New()
	}
	return res
}
