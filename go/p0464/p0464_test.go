// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0464

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	arg2   int
	target bool
}

var values = []result{
	{
		arg1:   10,
		arg2:   11,
		target: false,
	},
	{
		arg1:   11,
		arg2:   10,
		target: true,
	},
	{
		arg1:   10,
		arg2:   0,
		target: true,
	},
	{
		arg1:   1,
		arg2:   2,
		target: false,
	},
}

type p0464TestSuite struct {
	suite.Suite
}

func (s *p0464TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, canIWin(v.arg1, v.arg2))
	}
}

func TestP0464TestSuite(t *testing.T) {
	s := &p0464TestSuite{}
	suite.Run(t, s)
}
