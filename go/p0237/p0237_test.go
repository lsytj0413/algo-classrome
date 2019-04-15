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

package p0237

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"algo/comm"
)

type p0237TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	arg2   int
	target []int
}

var values = []*result{
	{
		arg1:   []int{4, 5, 1, 9},
		arg2:   5,
		target: []int{4, 1, 9},
	},
	{
		arg1:   []int{4, 5, 1, 9},
		arg2:   1,
		target: []int{4, 5, 9},
	},
}

func (s *p0237TestSuite) Test() {
	for _, v := range values {
		l := comm.Links(v.arg1)
		n := comm.FindListNode(l, v.arg2)
		deleteNode(n)

		comm.AssertLinkEqual(s.Suite, comm.Links(v.target), l)
	}
}

func TestP0237TestSuite(t *testing.T) {
	s := &p0237TestSuite{}
	suite.Run(t, s)
}
