// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0997

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	arg2   [][]int
	target int
}

var values = []result{
	{
		arg1: 2,
		arg2: [][]int{
			[]int{1, 2},
		},
		target: 2,
	},
	{
		arg1: 3,
		arg2: [][]int{
			[]int{1, 3},
			[]int{2, 3},
		},
		target: 3,
	},
	{
		arg1: 3,
		arg2: [][]int{
			[]int{1, 3},
			[]int{2, 3},
			[]int{3, 1},
		},
		target: -1,
	},
	{
		arg1: 3,
		arg2: [][]int{
			[]int{1, 2},
			[]int{2, 3},
		},
		target: -1,
	},
	{
		arg1: 4,
		arg2: [][]int{
			[]int{1, 3},
			[]int{1, 4},
			[]int{2, 3},
			[]int{2, 4},
			[]int{4, 3},
		},
		target: 3,
	},
}

type p0997TestSuite struct {
	suite.Suite
}

func (s *p0997TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findJudge(v.arg1, v.arg2))
	}
}

func TestP0997TestSuite(t *testing.T) {
	s := &p0997TestSuite{}
	suite.Run(t, s)
}
