// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0372

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
		arg1:   2,
		arg2:   []int{3},
		target: 8,
	},
	{
		arg1:   2,
		arg2:   []int{1, 0},
		target: 1024,
	},
}

type p0372TestSuite struct {
	suite.Suite
}

func (s *p0372TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, superPow(v.arg1, v.arg2))
	}
}

func TestP0372TestSuite(t *testing.T) {
	s := &p0372TestSuite{}
	suite.Run(t, s)
}
