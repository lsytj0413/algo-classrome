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

package p0402

import "container/list"

func removeKdigits(num string, k int) string {
	n, size := k, len(num)
	l := list.New()
	for i := 0; i < size; i++ {
		val := int(num[i] - '0')
		for l.Len() != 0 && n > 0 {
			e := l.Back()
			if val >= e.Value.(int) {
				break
			}
			l.Remove(e)
			n--
		}
		l.PushBack(val)
	}

	cnt := 0
	for l.Len() != 0 {
		e := l.Front()
		if e.Value.(int) != 0 {
			break
		}
		l.Remove(e)
		cnt++
	}

	r := ""
	for l.Len() != 0 && len(r) < size-k-cnt {
		e := l.Front()
		r += string('0' + byte(e.Value.(int)))
		l.Remove(e)
	}

	if r == "" {
		return "0"
	}
	return r
}
