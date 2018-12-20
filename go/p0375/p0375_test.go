// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0375

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	target int
}

var values = []result{
	{
		arg1:   1,
		target: 0,
	},
	{
		arg1:   2,
		target: 1,
	},
}

type p0375TestSuite struct {
	suite.Suite
}

func (s *p0375TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, getMoneyAmount(v.arg1))
	}
}

func TestP0375TestSuite(t *testing.T) {
	s := &p0375TestSuite{}
	suite.Run(t, s)
}
