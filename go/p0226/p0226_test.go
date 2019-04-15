// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0226

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
		arg1:   []int{4, 2, 7, 1, 3, 6, 9},
		target: []int{4, 7, 2, 9, 6, 3, 1},
	},
}

type p0226TestSuite struct {
	suite.Suite
}

func (s *p0226TestSuite) Test() {
	for _, v := range values {
		comm.AssertTreeEqual(s.Suite, comm.Trees(v.target), invertTree(comm.Trees(v.arg1)))
	}
}

func TestP0226TestSuite(t *testing.T) {
	s := &p0226TestSuite{}
	suite.Run(t, s)
}
