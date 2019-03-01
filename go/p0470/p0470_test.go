// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0470

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
		arg1:   []int{1, 2, 3, 4},
		target: 3,
	},
}

type p0470TestSuite struct {
	suite.Suite
}

func (s *p0470TestSuite) Test() {
	for _, v := range values {
		_ = v
		// s.Equal(v.target, numberOfArithmeticSlices(v.arg1))
	}
}

func TestP0470TestSuite(t *testing.T) {
	s := &p0470TestSuite{}
	suite.Run(t, s)
}
