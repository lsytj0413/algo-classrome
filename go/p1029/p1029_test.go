// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p1029

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
			[]int{10, 20},
			[]int{30, 200},
			[]int{400, 50},
			[]int{30, 20},
		},
		target: 110,
	},
}

type p1029TestSuite struct {
	suite.Suite
}

func (s *p1029TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, twoCitySchedCost(v.arg1))
	}
}

func TestP1029TestSuite(t *testing.T) {
	s := &p1029TestSuite{}
	suite.Run(t, s)
}
