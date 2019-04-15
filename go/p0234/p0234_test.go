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

package p0234

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"algo/comm"
)

type p0234TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	target bool
}

var values = []result{
	{
		arg1:   []int{1, 2},
		target: false,
	},
	{
		arg1:   []int{1, 2, 2, 1},
		target: true,
	},
}

func (s *p0234TestSuite) TestAddTwoNumbers() {
	for _, v := range values {
		s.Equal(v.target, isPalindrome(comm.Links(v.arg1)))
	}
}

func TestP0234TestSuite(t *testing.T) {
	s := &p0234TestSuite{}
	suite.Run(t, s)
}
