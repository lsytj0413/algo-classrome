// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0424

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   int
	target int
}

var values = []result{
	{
		arg1:   "ABAB",
		arg2:   2,
		target: 4,
	},
	{
		arg1:   "AABABBA",
		arg2:   1,
		target: 4,
	},
	{
		arg1:   "AAAA",
		arg2:   0,
		target: 4,
	},
}

type p0424TestSuite struct {
	suite.Suite
}

func (s *p0424TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, characterReplacement(v.arg1, v.arg2))
	}
}

func TestP0424TestSuite(t *testing.T) {
	s := &p0424TestSuite{}
	suite.Run(t, s)
}
