// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0897

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
		arg1:   []int{5, 3, 6},
		target: []int{3, 0xFFFFFFFF, 5, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 6},
	},
}

type p0897TestSuite struct {
	suite.Suite
}

func (s *p0897TestSuite) Test() {
	for _, v := range values {
		r := increasingBST(comm.Trees(v.arg1))
		comm.AssertTreeEqual(s.Suite, comm.Trees(v.target), r)
	}
}

func TestP0897TestSuite(t *testing.T) {
	s := &p0897TestSuite{}
	suite.Run(t, s)
}
