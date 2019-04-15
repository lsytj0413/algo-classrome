// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0206

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"algo/comm"
)

type result struct {
	arg1   []int
	target []int
}

var values = []result{
	{
		arg1:   []int{1, 2, 3, 4, 5},
		target: []int{5, 4, 3, 2, 1},
	},
}

type p0206TestSuite struct {
	suite.Suite
}

func (s *p0206TestSuite) Test() {
	for _, v := range values {
		comm.AssertLinkEqual(s.Suite, comm.Links(v.target), reverseList(comm.Links(v.arg1)))
	}
}

func TestP0206TestSuite(t *testing.T) {
	s := &p0206TestSuite{}
	suite.Run(t, s)
}
