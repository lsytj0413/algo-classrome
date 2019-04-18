// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0948

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
		arg1:   []int{100},
		arg2:   50,
		target: 0,
	},
	{
		arg1:   []int{100, 200},
		arg2:   150,
		target: 1,
	},
	{
		arg1:   []int{100, 200, 300, 400},
		arg2:   200,
		target: 2,
	},
}

type p0948TestSuite struct {
	suite.Suite
}

func (s *p0948TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, bagOfTokensScore(v.arg1, v.arg2))
	}
}

func TestP0948TestSuite(t *testing.T) {
	s := &p0948TestSuite{}
	suite.Run(t, s)
}
