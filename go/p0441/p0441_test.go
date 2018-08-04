// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0441

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
		arg1:   5,
		target: 2,
	},
	{
		arg1:   8,
		target: 3,
	},
}

type p0441TestSuite struct {
	suite.Suite
}

func (s *p0441TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, arrangeCoins(v.arg1))
	}
}

func TestP0441TestSuite(t *testing.T) {
	s := &p0441TestSuite{}
	suite.Run(t, s)
}
