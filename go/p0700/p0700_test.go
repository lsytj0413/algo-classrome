// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0700

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"algo/comm"
)

type result struct {
	arg1   []int
	arg2   int
	target []int
}

var values = []result{
	{
		arg1:   []int{4, 2, 7, 1, 3},
		arg2:   2,
		target: []int{2, 1, 3},
	},
	{
		arg1:   []int{4, 2, 7, 1, 3},
		arg2:   5,
		target: []int{},
	},
}

type p0700TestSuite struct {
	suite.Suite
}

func (s *p0700TestSuite) Test() {
	for _, v := range values {
		r := searchBST(comm.Trees(v.arg1), v.arg2)
		comm.AssertTreeEqual(s.Suite, comm.Trees(v.target), r)
	}
}

func TestP0700TestSuite(t *testing.T) {
	s := &p0700TestSuite{}
	suite.Run(t, s)
}
