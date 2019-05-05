// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p1005

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
		arg1:   []int{4, 2, 3},
		arg2:   1,
		target: 5,
	},
	{
		arg1:   []int{3, -1, 0, 2},
		arg2:   3,
		target: 6,
	},
	{
		arg1:   []int{2, -3, -1, 5, -4},
		arg2:   2,
		target: 13,
	},
}

type p1005TestSuite struct {
	suite.Suite
}

func (s *p1005TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, largestSumAfterKNegations(v.arg1, v.arg2))
	}
}

func TestP1005TestSuite(t *testing.T) {
	s := &p1005TestSuite{}
	suite.Run(t, s)
}
