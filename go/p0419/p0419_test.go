// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0419

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   [][]byte
	target int
}

var values = []result{
	{
		arg1: [][]byte{
			[]byte("X..X"),
			[]byte("...X"),
			[]byte("...X"),
		},
		target: 2,
	},
}

type p0419TestSuite struct {
	suite.Suite
}

func (s *p0419TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, countBattleships(v.arg1))
	}
}

func TestP0419TestSuite(t *testing.T) {
	s := &p0419TestSuite{}
	suite.Run(t, s)
}
