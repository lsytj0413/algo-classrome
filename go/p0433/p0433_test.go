// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0433

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   string
	arg3   []string
	target int
}

var values = []result{
	// {
	// 	arg1: "AACCGGTT",
	// 	arg2: "AACCGGTA",
	// 	arg3: []string{
	// 		"AACCGGTA",
	// 	},
	// 	target: 1,
	// },
	// {
	// 	arg1: "AACCGGTT",
	// 	arg2: "AAACGGTA",
	// 	arg3: []string{
	// 		"AACCGGTA",
	// 		"AACCGCTA",
	// 		"AAACGGTA",
	// 	},
	// 	target: 2,
	// },
	// {
	// 	arg1: "AAAAACCC",
	// 	arg2: "AACCCCCC",
	// 	arg3: []string{
	// 		"AAAACCCC",
	// 		"AAACCCCC",
	// 		"AACCCCCC",
	// 	},
	// 	target: 3,
	// },
	{
		arg1: "AAAACCCC",
		arg2: "CCCCCCCC",
		arg3: []string{
			"AAAACCCA", "AAACCCCA", "AACCCCCA", "AACCCCCC", "ACCCCCCC", "CCCCCCCC", "AAACCCCC", "AACCCCCC",
		},
		target: 4,
	},
}

type p0433TestSuite struct {
	suite.Suite
}

func (s *p0433TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, minMutation(v.arg1, v.arg2, v.arg3))
	}
}

func TestP0433TestSuite(t *testing.T) {
	s := &p0433TestSuite{}
	suite.Run(t, s)
}
