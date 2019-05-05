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

package p0933

import (
	"container/list"
)

type RecentCounter struct {
	list *list.List
}

func Constructor() RecentCounter {
	return RecentCounter{
		list: list.New(),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.list.PushBack(t)

	for this.list.Len() > 0 {
		e := this.list.Front()
		v := e.Value.(int)
		if v < t-3000 {
			this.list.Remove(e)
		} else {
			break
		}
	}

	return this.list.Len()
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
