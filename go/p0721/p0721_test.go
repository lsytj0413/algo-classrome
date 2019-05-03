// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0721

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   [][]string
	target [][]string
}

var values = []result{
	{
		arg1: [][]string{
			[]string{"John", "johnsmith@mail.com", "john00@mail.com"},
			[]string{"John", "johnnybravo@mail.com"},
			[]string{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
			[]string{"Mary", "mary@mail.com"},
		},
		target: [][]string{
			[]string{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
			[]string{"John", "johnnybravo@mail.com"},
			[]string{"Mary", "mary@mail.com"},
		},
	},
	{
		arg1: [][]string{
			[]string{"David", "David0@m.co", "David1@m.co"},
			[]string{"David", "David3@m.co", "David4@m.co"},
			[]string{"David", "David4@m.co", "David5@m.co"},
			[]string{"David", "David2@m.co", "David3@m.co"},
			[]string{"David", "David1@m.co", "David2@m.co"},
		},
		target: [][]string{
			[]string{"David", "David0@m.co", "David1@m.co", "David2@m.co", "David3@m.co", "David4@m.co", "David5@m.co"},
		},
	},
}

type p0721TestSuite struct {
	suite.Suite
}

func (s *p0721TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, accountsMerge(v.arg1))
	}
}

func TestP0721TestSuite(t *testing.T) {
	s := &p0721TestSuite{}
	suite.Run(t, s)
}
