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

package p0002

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"algo/comm"
)

type p0002TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	arg2   []int
	target []int
}

var values = []result{
	{
		arg1:   []int{2, 4, 3},
		arg2:   []int{5, 6, 4},
		target: []int{7, 0, 8},
	},
	{
		arg1:   []int{0},
		arg2:   []int{0},
		target: []int{0},
	},
	{
		arg1:   []int{0},
		arg2:   []int{1},
		target: []int{1},
	},
	{
		arg1:   []int{2},
		arg2:   []int{0},
		target: []int{2},
	},
	{
		arg1:   []int{1},
		arg2:   []int{1},
		target: []int{2},
	},
	{
		arg1:   []int{5},
		arg2:   []int{3},
		target: []int{8},
	},
	{
		arg1:   []int{4},
		arg2:   []int{5},
		target: []int{9},
	},
	{
		arg1:   []int{5},
		arg2:   []int{5},
		target: []int{0, 1},
	},
	{
		arg1:   []int{9},
		arg2:   []int{9},
		target: []int{8, 1},
	},
	{
		arg1:   []int{6},
		arg2:   []int{8},
		target: []int{4, 1},
	},
	{
		arg1:   []int{1, 8},
		arg2:   []int{0},
		target: []int{1, 8},
	},
	{
		arg1:   []int{0},
		arg2:   []int{7, 3},
		target: []int{7, 3},
	},
	{
		arg1:   []int{9, 8},
		arg2:   []int{1},
		target: []int{0, 9},
	},
	{
		arg1:   []int{1},
		arg2:   []int{9, 9},
		target: []int{0, 0, 1},
	},
	{
		arg1:   []int{9, 9},
		arg2:   []int{9},
		target: []int{8, 0, 1},
	},
	{
		arg1:   []int{9, 1, 6},
		arg2:   []int{0},
		target: []int{9, 1, 6},
	},
	{
		arg1:   []int{0},
		arg2:   []int{2, 7, 8},
		target: []int{2, 7, 8},
	},
	{
		arg1:   []int{8, 9, 9},
		arg2:   []int{2},
		target: []int{0, 0, 0, 1},
	},
	{
		arg1:   []int{9},
		arg2:   []int{9, 9, 9},
		target: []int{8, 0, 0, 1},
	},
	{
		arg1:   []int{3, 7},
		arg2:   []int{9, 2},
		target: []int{2, 0, 1},
	},
	{
		arg1:   []int{9, 9},
		arg2:   []int{9, 9},
		target: []int{8, 9, 1},
	},
	{
		arg1:   []int{1, 8, 3},
		arg2:   []int{7, 1},
		target: []int{8, 9, 3},
	},
	{
		arg1:   []int{8, 6},
		arg2:   []int{6, 4, 8},
		target: []int{4, 1, 9},
	},
	{
		arg1:   []int{8, 6},
		arg2:   []int{6, 4, 9},
		target: []int{4, 1, 0, 1},
	},
	{
		arg1:   []int{2, 8, 9},
		arg2:   []int{5, 1, 9},
		target: []int{7, 9, 8, 1},
	},
	{
		arg1:   []int{1, 0, 9},
		arg2:   []int{5, 7, 8},
		target: []int{6, 7, 7, 1},
	},
	{
		arg1:   []int{9, 9, 9, 9, 9},
		arg2:   []int{1},
		target: []int{0, 0, 0, 0, 0, 1},
	},
	{
		arg1:   []int{9},
		arg2:   []int{1, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		target: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	},
	{
		arg1:   []int{3, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		arg2:   []int{7},
		target: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	},
	{
		arg1:   []int{9},
		arg2:   []int{1, 9, 9, 9, 9, 9, 8, 9, 9, 9},
		target: []int{0, 0, 0, 0, 0, 0, 9, 9, 9, 9},
	},
	{
		arg1:   []int{2, 8, 9, 9, 9, 9, 8, 9, 9, 9},
		arg2:   []int{8},
		target: []int{0, 9, 9, 9, 9, 9, 8, 9, 9, 9},
	},
	{
		arg1:   []int{2, 8, 9, 9, 9, 9, 8, 9, 9, 9},
		arg2:   []int{8, 1, 2},
		target: []int{0, 0, 2, 0, 0, 0, 9, 9, 9, 9},
	},
	{
		arg1:   []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		arg2:   []int{0},
		target: []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
	},
	{
		arg1:   []int{0},
		arg2:   []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		target: []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
	},
	{
		arg1:   []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		arg2:   []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		target: []int{8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 1},
	},
	{
		arg1:   []int{8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		arg2:   []int{8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		target: []int{6, 7, 7, 7, 7, 7, 7, 7, 7, 7, 1},
	},
	{
		arg1:   []int{8, 3, 2, 7, 4, 5, 7, 9, 8, 1},
		arg2:   []int{2, 6, 7, 2, 5, 4, 2, 0, 1, 8},
		target: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	},
	{
		arg1:   []int{3, 1, 0, 0, 1, 9, 0, 1, 6, 1},
		arg2:   []int{5, 5, 8, 6, 2, 5, 8, 2, 6, 1},
		target: []int{8, 6, 8, 6, 3, 4, 9, 3, 2, 3},
	},
	{
		arg1:   []int{6, 4, 5, 0, 4, 4, 9, 4, 1},
		arg2:   []int{3, 8, 8, 0, 3, 0, 1, 4, 8},
		target: []int{9, 2, 4, 1, 7, 4, 0, 9, 9},
	},
	{
		arg1:   []int{7, 6, 8, 4, 1, 6, 5, 9, 6, 1},
		arg2:   []int{6, 4, 9, 1, 9, 5, 6, 5, 0, 1},
		target: []int{3, 1, 8, 6, 0, 2, 2, 5, 7, 2},
	},
	{
		arg1:   []int{1, 6, 0, 3, 3, 6, 7, 2, 0, 1},
		arg2:   []int{6, 3, 0, 8, 9, 6, 6, 9, 6, 1},
		target: []int{7, 9, 0, 1, 3, 3, 4, 2, 7, 2},
	},
	{
		arg1:   []int{9, 1, 6, 9, 9, 5, 8, 7, 1, 1},
		arg2:   []int{6, 3, 6, 8, 0, 6, 3, 1, 1, 1},
		target: []int{5, 5, 2, 8, 0, 2, 2, 9, 2, 2},
	},
	{
		arg1:   []int{8, 6, 5, 8, 4, 5, 9, 3, 4},
		arg2:   []int{3, 4, 3, 4, 1, 8, 9, 5, 5},
		target: []int{1, 1, 9, 2, 6, 3, 9, 9, 9},
	},
	{
		arg1:   []int{3, 2, 5, 9, 0, 5, 0, 1, 5, 1},
		arg2:   []int{7, 7, 8, 5, 5, 3, 4, 9, 8},
		target: []int{0, 0, 4, 5, 6, 8, 4, 0, 4, 2},
	},
}

func (s *p0002TestSuite) Test() {
	for _, v := range values {
		actual := addTwoNumbers(comm.Links(v.arg1), comm.Links(v.arg2))
		comm.AssertLinkEqual(s.Suite, comm.Links(v.target), actual)
	}
}

func TestP0002TestSuite(t *testing.T) {
	s := &p0002TestSuite{}
	suite.Run(t, s)
}
