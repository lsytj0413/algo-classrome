// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p1011

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   int
	target int
}

var values = []result{
	{
		arg1:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		arg2:   5,
		target: 15,
	},
	{
		arg1:   []int{3, 2, 2, 4, 1, 4},
		arg2:   3,
		target: 6,
	},
	{
		arg1:   []int{1, 2, 3, 1, 1},
		arg2:   4,
		target: 3,
	},
}

type p1011TestSuite struct {
	suite.Suite
}

func (s *p1011TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, shipWithinDays(v.arg1, v.arg2))
	}
}

func TestP1011TestSuite(t *testing.T) {
	s := &p1011TestSuite{}
	suite.Run(t, s)
}
