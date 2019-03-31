// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0529

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   [][]byte
	arg2   []int
	target [][]byte
}

var values = []result{
	{
		arg1: [][]byte{
			[]byte("EEEEE"),
			[]byte("EEMEE"),
			[]byte("EEEEE"),
			[]byte("EEEEE"),
		},
		arg2: []int{3, 0},
		target: [][]byte{
			[]byte("B1E1B"),
			[]byte("B1M1B"),
			[]byte("B111B"),
			[]byte("BBBBB"),
		},
	},
	{
		arg1: [][]byte{
			[]byte("B1E1B"),
			[]byte("B1M1B"),
			[]byte("B111B"),
			[]byte("BBBBB"),
		},
		arg2: []int{1, 2},
		target: [][]byte{
			[]byte("B1E1B"),
			[]byte("B1X1B"),
			[]byte("B111B"),
			[]byte("BBBBB"),
		},
	},
}

type p0529TestSuite struct {
	suite.Suite
}

func (s *p0529TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, updateBoard(v.arg1, v.arg2))
	}
}

func TestP0529TestSuite(t *testing.T) {
	s := &p0529TestSuite{}
	suite.Run(t, s)
}
