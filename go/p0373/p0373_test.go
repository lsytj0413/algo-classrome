// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0373

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   []int
	arg3   int
	target [][]int
}

var values = []result{
	{
		arg1: []int{1, 7, 11},
		arg2: []int{2, 4, 6},
		arg3: 3,
		target: [][]int{
			[]int{1, 2},
			[]int{1, 4},
			[]int{1, 6},
		},
	},
	{
		arg1: []int{1, 1, 2},
		arg2: []int{1, 2, 3},
		arg3: 2,
		target: [][]int{
			[]int{1, 1},
			[]int{1, 1},
		},
	},
	{
		arg1: []int{1, 2},
		arg2: []int{3},
		arg3: 3,
		target: [][]int{
			[]int{1, 3},
			[]int{2, 3},
		},
	},
	{
		arg1: []int{1, 1, 2},
		arg2: []int{1, 2, 3},
		arg3: 10,
		target: [][]int{
			[]int{1, 1},
			[]int{1, 1},
			[]int{2, 1},
			[]int{1, 2},
			[]int{1, 2},
			[]int{2, 2},
			[]int{1, 3},
			[]int{1, 3},
			[]int{2, 3},
		},
	},
}

type p0373TestSuite struct {
	suite.Suite
}

func (s *p0373TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, kSmallestPairs(v.arg1, v.arg2, v.arg3))
	}
}

func TestP0373TestSuite(t *testing.T) {
	s := &p0373TestSuite{}
	suite.Run(t, s)
}
