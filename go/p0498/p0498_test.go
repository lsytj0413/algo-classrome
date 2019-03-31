// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0498

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   [][]int
	target []int
}

var values = []result{
	{
		arg1: [][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		},
		target: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
	},
}

type p0498TestSuite struct {
	suite.Suite
}

func (s *p0498TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findDiagonalOrder(v.arg1))
	}
}

func TestP0498TestSuite(t *testing.T) {
	s := &p0498TestSuite{}
	suite.Run(t, s)
}
