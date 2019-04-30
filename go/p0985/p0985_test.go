// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0985

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   [][]int
	target []int
}

var values = []result{
	{
		arg1: []int{1, 2, 3, 4},
		arg2: [][]int{
			[]int{1, 0},
			[]int{-3, 1},
			[]int{-4, 0},
			[]int{2, 3},
		},
		target: []int{8, 6, 2, 4},
	},
}

type p0985TestSuite struct {
	suite.Suite
}

func (s *p0985TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, sumEvenAfterQueries(v.arg1, v.arg2))
	}
}

func TestP0985TestSuite(t *testing.T) {
	s := &p0985TestSuite{}
	suite.Run(t, s)
}
