// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p1033

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	arg2   int
	arg3   int
	target []int
}

var values = []result{
	{
		arg1:   1,
		arg2:   2,
		arg3:   5,
		target: []int{1, 2},
	},
	{
		arg1:   4,
		arg2:   3,
		arg3:   2,
		target: []int{0, 0},
	},
	{
		arg1:   3,
		arg2:   5,
		arg3:   1,
		target: []int{1, 2},
	},
}

type p1033TestSuite struct {
	suite.Suite
}

func (s *p1033TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, numMovesStones(v.arg1, v.arg2, v.arg3))
	}
}

func TestP1033TestSuite(t *testing.T) {
	s := &p1033TestSuite{}
	suite.Run(t, s)
}
