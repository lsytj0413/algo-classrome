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

package p0083

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"algo/comm"
)

type p0083TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	target []int
}

var values = []result{
	{
		arg1:   []int{1, 1, 2},
		target: []int{1, 2},
	},
	{
		arg1:   []int{1, 1, 2, 3, 3},
		target: []int{1, 2, 3},
	},
}

func (s *p0083TestSuite) TestAddTwoNumbers() {
	for _, v := range values {
		actual := deleteDuplicates(comm.Links(v.arg1))
		comm.AssertLinkEqual(s.Suite, comm.Links(v.target), actual)
	}
}

func TestP0083TestSuite(t *testing.T) {
	s := &p0083TestSuite{}
	suite.Run(t, s)
}
