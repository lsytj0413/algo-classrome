// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0021

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"algo/comm"
)

type result struct {
	arg1   []int
	arg2   []int
	target []int
}

var values = []result{
	{
		arg1:   []int{1, 2, 4},
		arg2:   []int{1, 3, 4},
		target: []int{1, 1, 2, 3, 4, 4},
	},
}

type p0021TestSuite struct {
	suite.Suite
}

func (s *p0021TestSuite) Test() {
	for _, v := range values {
		l := mergeTwoLists(comm.Links(v.arg1), comm.Links(v.arg2))
		comm.AssertLinkEqual(s.Suite, comm.Links(v.target), l)
	}
}

func TestP0021TestSuite(t *testing.T) {
	s := &p0021TestSuite{}
	suite.Run(t, s)
}
