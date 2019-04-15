// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0617

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
		arg1:   []int{1, 3, 2, 5},
		arg2:   []int{2, 1, 3, 0xFFFFFFFF, 4, 0xFFFFFFFF, 7},
		target: []int{3, 4, 5, 5, 4, 0xFFFFFFFF, 7},
	},
}

type p0617TestSuite struct {
	suite.Suite
}

func (s *p0617TestSuite) Test() {
	for _, v := range values {
		r := mergeTrees(comm.Trees(v.arg1), comm.Trees(v.arg2))
		comm.AssertTreeEqual(s.Suite, comm.Trees(v.target), r)
	}
}

func TestP0617TestSuite(t *testing.T) {
	s := &p0617TestSuite{}
	suite.Run(t, s)
}
