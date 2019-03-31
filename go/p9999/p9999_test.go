// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p9999

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   [][]int
	target int
}

var values = []result{
	{
		arg1: [][]int{
			[]int{3, 0, -2, 4, 0},
			[]int{-1, 2, -2, 1, 4},
			[]int{3, 1, -2, -3, 3},
			[]int{2, -4, -3, -3, 2},
			[]int{5, 2, -2, -3, 1},
		},
		target: -12,
	},
}

type p9999TestSuite struct {
	suite.Suite
}

func (s *p9999TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, minValue(v.arg1))
	}
}

func TestP9999TestSuite(t *testing.T) {
	s := &p9999TestSuite{}
	suite.Run(t, s)
}
