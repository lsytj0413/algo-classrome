// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0481

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
		arg1:   6,
		target: 3,
	},
}

type p0481TestSuite struct {
	suite.Suite
}

func (s *p0481TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, magicalString(v.arg1))
	}
}

func TestP0481TestSuite(t *testing.T) {
	s := &p0481TestSuite{}
	suite.Run(t, s)
}
