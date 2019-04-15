// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0653

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
		arg1:   []int{5, 3, 6, 2, 4, 0xFFFFFFFF, 7},
		arg2:   9,
		target: true,
	},
	{
		arg1:   []int{5, 3, 6, 2, 4, 0xFFFFFFFF, 7},
		arg2:   28,
		target: false,
	},
}

type p0653TestSuite struct {
	suite.Suite
}

func (s *p0653TestSuite) Test() {
	for _, v := range values {
		r := findTarget(comm.Trees(v.arg1), v.arg2)
		s.Equal(v.target, r)
	}
}

func TestP0653TestSuite(t *testing.T) {
	s := &p0653TestSuite{}
	suite.Run(t, s)
}
