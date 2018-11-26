// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0378

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   [][]int
	arg2   int
	target int
}

var values = []result{
	{
		arg1: [][]int{
			[]int{1, 5, 9},
			[]int{10, 11, 13},
			[]int{12, 13, 15},
		},
		arg2:   8,
		target: 13,
	},
}

type p0378TestSuite struct {
	suite.Suite
}

func (s *p0378TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, kthSmallest(v.arg1, v.arg2))
	}
}

func TestP0378TestSuite(t *testing.T) {
	s := &p0378TestSuite{}
	suite.Run(t, s)
}
