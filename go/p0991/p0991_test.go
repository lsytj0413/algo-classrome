// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0991

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	arg2   int
	target int
}

var values = []result{
	{
		arg1:   2,
		arg2:   3,
		target: 2,
	},
	{
		arg1:   5,
		arg2:   8,
		target: 2,
	},
	{
		arg1:   3,
		arg2:   10,
		target: 3,
	},
	{
		arg1:   1024,
		arg2:   1,
		target: 1023,
	},
	{
		arg1:   1,
		arg2:   1000000000,
		target: 39,
	},
}

type p0991TestSuite struct {
	suite.Suite
}

func (s *p0991TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, brokenCalc(v.arg1, v.arg2))
	}
}

func TestP0991TestSuite(t *testing.T) {
	s := &p0991TestSuite{}
	suite.Run(t, s)
}
