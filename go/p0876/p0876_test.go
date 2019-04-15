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

package p0876

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"algo/comm"
)

type p0876TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	target []int
}

var values = []result{
	{
		arg1:   []int{1, 2, 3, 4, 5},
		target: []int{3, 4, 5},
	},
	{
		arg1:   []int{1, 2, 3, 4, 5, 6},
		target: []int{4, 5, 6},
	},
}

func (s *p0876TestSuite) Test() {
	for _, v := range values {
		actual := middleNode(comm.Links(v.arg1))
		comm.AssertLinkEqual(s.Suite, comm.Links(v.target), actual)
	}
}

func TestP0876TestSuite(t *testing.T) {
	s := &p0876TestSuite{}
	suite.Run(t, s)
}
