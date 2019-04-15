// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0437

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"algo/comm"
)

type result struct {
	arg1   []int
	arg2   int
	target int
}

var values = []result{
	{
		arg1:   []int{10, 5, -3, 3, 2, 0xFFFFFFFF, 11, 3, -2, 0xFFFFFFFF, 1},
		arg2:   8,
		target: 3,
	},
}

type p0437TestSuite struct {
	suite.Suite
}

func (s *p0437TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, pathSum(comm.Trees(v.arg1), v.arg2))
	}
}

func TestP0437TestSuite(t *testing.T) {
	s := &p0437TestSuite{}
	suite.Run(t, s)
}
