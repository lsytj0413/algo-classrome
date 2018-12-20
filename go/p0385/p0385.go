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

package p0385

import (
	"container/list"
	"strconv"
)

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

// NestedInteger ...
type NestedInteger struct {
	isInteger bool
	v         int
	sub       []*NestedInteger
}

func (n NestedInteger) IsInteger() bool {
	return n.isInteger
}

func (n NestedInteger) GetInteger() int {
	return n.v
}

func (n *NestedInteger) SetInteger(value int) {
	n.isInteger = true
	n.v = value
}

func (n *NestedInteger) Add(elem NestedInteger) {}

func (n *NestedInteger) GetList() []*NestedInteger {
	return nil
}

func deserialize(s string) *NestedInteger {
	atoi := func(s string) int {
		v, _ := strconv.Atoi(s)
		return v
	}

	nestedInt := func(s string) *NestedInteger {
		n := &NestedInteger{}
		n.SetInteger(atoi(s))
		return n
	}

	if s[0] != '[' {
		return nestedInt(s)
	}

	stack := list.New()
	var ans *NestedInteger
	idx := 0

	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == '[':
			n := &NestedInteger{}
			stack.PushBack(n)
			if ans == nil {
				ans = n
			}
			idx = i + 1
		case s[i] == ',':
			if idx != i {
				top := stack.Back().Value.(*NestedInteger)
				top.Add(*nestedInt(s[idx:i]))
			}
			idx = i + 1
		case s[i] == ']':
			if idx != i {
				top := stack.Back().Value.(*NestedInteger)
				top.Add(*nestedInt(s[idx:i]))
			}
			cur := stack.Back().Value.(*NestedInteger)
			e := stack.Back()
			stack.Remove(e)
			if stack.Len() != 0 {
				top := stack.Back().Value.(*NestedInteger)
				top.Add(*cur)
			}
			idx = i + 1
		}
	}

	return ans
}
