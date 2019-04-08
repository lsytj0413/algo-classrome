// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0898

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
		arg1:   []int{0},
		target: 1,
	},
	{
		arg1:   []int{1, 1, 2},
		target: 3,
	},
	{
		arg1:   []int{1, 2, 4},
		target: 6,
	},
	{
		arg1:   []int{8, 10, 9, 13},
		target: 6,
	},
}

type p0898TestSuite struct {
	suite.Suite
}

func (s *p0898TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, subarrayBitwiseORs(v.arg1))
	}
}

func TestP0898TestSuite(t *testing.T) {
	s := &p0898TestSuite{}
	suite.Run(t, s)
}
