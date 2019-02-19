// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0454

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   []int
	arg3   []int
	arg4   []int
	target int
}

var values = []result{
	{
		arg1:   []int{1, 2},
		arg2:   []int{-2, -1},
		arg3:   []int{-1, 2},
		arg4:   []int{0, 2},
		target: 2,
	},
}

type p0454TestSuite struct {
	suite.Suite
}

func (s *p0454TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, fourSumCount(v.arg1, v.arg2, v.arg3, v.arg4))
	}
}

func TestP0454TestSuite(t *testing.T) {
	s := &p0454TestSuite{}
	suite.Run(t, s)
}
