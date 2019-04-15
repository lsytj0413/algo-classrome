// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0100

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"algo/comm"
)

type result struct {
	arg1   []int
	arg2   []int
	target bool
}

var values = []result{
	{
		arg1:   []int{1, 2, 3},
		arg2:   []int{1, 2, 3},
		target: true,
	},
	{
		arg1:   []int{1, 2},
		arg2:   []int{1, 0xFFFFFFFF, 2},
		target: false,
	},
	{
		arg1:   []int{1, 2, 1},
		arg2:   []int{1, 1, 2},
		target: false,
	},
}

type p0100TestSuite struct {
	suite.Suite
}

func (s *p0100TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, isSameTree(comm.Trees(v.arg1), comm.Trees(v.arg2)))
	}
}

func TestP0100TestSuite(t *testing.T) {
	s := &p0100TestSuite{}
	suite.Run(t, s)
}
