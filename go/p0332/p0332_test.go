// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0332

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   [][]string
	target []string
}

var values = []result{
	{
		arg1: [][]string{
			[]string{"MUC", "LHR"},
			[]string{"JFK", "MUC"},
			[]string{"SFO", "SJC"},
			[]string{"LHR", "SFO"},
		},
		target: []string{"JFK", "MUC", "LHR", "SFO", "SJC"},
	},
	{
		arg1: [][]string{
			[]string{"JFK", "SFO"},
			[]string{"JFK", "ATL"},
			[]string{"SFO", "ATL"},
			[]string{"ATL", "JFK"},
			[]string{"ATL", "SFO"},
		},
		target: []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
	},
}

type p0332TestSuite struct {
	suite.Suite
}

func (s *p0332TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findItinerary(v.arg1))
	}
}

func TestP0332TestSuite(t *testing.T) {
	s := &p0332TestSuite{}
	suite.Run(t, s)
}
