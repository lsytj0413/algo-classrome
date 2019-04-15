// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0112

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"algo/comm"
)

type result struct {
	arg1   []int
	arg2   int
	target bool
}

var values = []result{
	{
		arg1:   []int{5, 4, 8, 11, 0xFFFFFFFF, 13, 4, 7, 2, 0xFFFFFFFF, 1},
		arg2:   22,
		target: true,
	},
	{
		arg1:   []int{},
		arg2:   0,
		target: false,
	},
}

type p0112TestSuite struct {
	suite.Suite
}

func (s *p0112TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, hasPathSum(comm.Trees(v.arg1), v.arg2))
	}
}

func TestP0112TestSuite(t *testing.T) {
	s := &p0112TestSuite{}
	suite.Run(t, s)
}
