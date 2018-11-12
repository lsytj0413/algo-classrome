// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0338

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	target []int
}

var values = []result{
	{
		arg1:   2,
		target: []int{0, 1, 1},
	},
	{
		arg1:   5,
		target: []int{0, 1, 1, 2, 1, 2},
	},
}

type p0338TestSuite struct {
	suite.Suite
}

func (s *p0338TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, countBits(v.arg1))
	}
}

func TestP0338TestSuite(t *testing.T) {
	s := &p0338TestSuite{}
	suite.Run(t, s)
}
