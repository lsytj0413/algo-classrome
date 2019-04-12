// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0929

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []string
	target int
}

var values = []result{
	{
		arg1:   []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"},
		target: 2,
	},
}

type p0929TestSuite struct {
	suite.Suite
}

func (s *p0929TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, numUniqueEmails(v.arg1))
	}
}

func TestP0929TestSuite(t *testing.T) {
	s := &p0929TestSuite{}
	suite.Run(t, s)
}
