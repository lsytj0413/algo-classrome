// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0399

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   [][]string
	arg2   []float64
	arg3   [][]string
	target []float64
}

var values = []result{
	{
		arg1: [][]string{
			[]string{"a", "b"},
			[]string{"b", "c"},
		},
		arg2: []float64{2.0, 3.0},
		arg3: [][]string{
			[]string{"a", "c"},
			[]string{"b", "a"},
			[]string{"a", "e"},
			[]string{"a", "a"},
			[]string{"x", "x"},
		},
		target: []float64{6.0, 0.5, -1.0, 1.0, -1.0},
	},
}

type p0399TestSuite struct {
	suite.Suite
}

func (s *p0399TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, calcEquation(v.arg1, v.arg2, v.arg3))
	}
}

func TestP0399TestSuite(t *testing.T) {
	s := &p0399TestSuite{}
	suite.Run(t, s)
}
