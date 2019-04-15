// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0235

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"algo/comm"
)

type result struct {
	arg1   []int
	arg2   []int
	arg3   []int
	target int
}

var values = []result{
	{
		arg1:   []int{6, 2, 8, 0, 4, 7, 9, 0xFFFFFFFF, 0xFFFFFFFF, 3, 5},
		arg2:   []int{2},
		arg3:   []int{8},
		target: 6,
	},
	{
		arg1:   []int{6, 2, 8, 0, 4, 7, 9, 0xFFFFFFFF, 0xFFFFFFFF, 3, 5},
		arg2:   []int{2},
		arg3:   []int{4},
		target: 2,
	},
}

type p0235TestSuite struct {
	suite.Suite
}

func (s *p0235TestSuite) Test() {
	for _, v := range values {
		n := lowestCommonAncestor(comm.Trees(v.arg1), comm.Trees(v.arg2), comm.Trees(v.arg3))
		s.Equal(v.target, n.Val)
	}
}

func TestP0235TestSuite(t *testing.T) {
	s := &p0235TestSuite{}
	suite.Run(t, s)
}
