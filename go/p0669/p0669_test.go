// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0669

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"algo/comm"
)

type result struct {
	arg1   []int
	arg2   int
	arg3   int
	target []int
}

var values = []result{
	{
		arg1:   []int{1, 0, 2},
		arg2:   1,
		arg3:   2,
		target: []int{1, 0xFFFFFFFF, 2},
	},
	{
		arg1:   []int{3, 0, 4, 0xFFFFFFFF, 2, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 1},
		arg2:   1,
		arg3:   3,
		target: []int{3, 2, 0xFFFFFFFF, 1},
	},
}

type p0669TestSuite struct {
	suite.Suite
}

func (s *p0669TestSuite) Test() {
	for _, v := range values {
		r := trimBST(comm.Trees(v.arg1), v.arg2, v.arg3)
		comm.AssertTreeEqual(s.Suite, comm.Trees(v.target), r)
	}
}

func TestP0669TestSuite(t *testing.T) {
	s := &p0669TestSuite{}
	suite.Run(t, s)
}
