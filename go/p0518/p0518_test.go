// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0518

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	arg2   []int
	target int
}

var values = []result{
	{
		arg1:   5,
		arg2:   []int{1, 2, 5},
		target: 4,
	},
	{
		arg1:   3,
		arg2:   []int{2},
		target: 0,
	},
}

type p0518TestSuite struct {
	suite.Suite
}

func (s *p0518TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, change(v.arg1, v.arg2))
	}
}

func TestP0518TestSuite(t *testing.T) {
	s := &p0518TestSuite{}
	suite.Run(t, s)
}
