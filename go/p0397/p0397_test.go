// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0397

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
		arg1:   8,
		target: 3,
	},
	{
		arg1:   7,
		target: 4,
	},
	{
		arg1:   10000000,
		target: 30,
	},
	{
		arg1:   3,
		target: 2,
	},
}

type p0397TestSuite struct {
	suite.Suite
}

func (s *p0397TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, integerReplacement(v.arg1))
	}
}

func TestP0397TestSuite(t *testing.T) {
	s := &p0397TestSuite{}
	suite.Run(t, s)
}
