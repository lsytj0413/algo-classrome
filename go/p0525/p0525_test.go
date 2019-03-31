// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0525

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	target int
}

var values = []result{
	{
		arg1:   []int{0, 1},
		target: 2,
	},
	{
		arg1:   []int{0, 1, 0},
		target: 2,
	},
	{
		arg1:   []int{0, 1, 0, 1},
		target: 4,
	},
	{
		arg1:   []int{0, 0, 1, 0, 0, 0, 1, 1},
		target: 6,
	},
}

type p0525TestSuite struct {
	suite.Suite
}

func (s *p0525TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findMaxLength(v.arg1))
	}
}

func TestP0525TestSuite(t *testing.T) {
	s := &p0525TestSuite{}
	suite.Run(t, s)
}
