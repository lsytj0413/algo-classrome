// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0970

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	arg2   int
	arg3   int
	target []int
}

var values = []result{
	{
		arg1:   2,
		arg2:   3,
		arg3:   10,
		target: []int{2, 3, 4, 5, 7, 9, 10},
	},
	{
		arg1:   3,
		arg2:   5,
		arg3:   15,
		target: []int{2, 4, 6, 8, 10, 14},
	},
	{
		arg1:   2,
		arg2:   1,
		arg3:   10,
		target: []int{2, 3, 5, 9},
	},
}

type p0970TestSuite struct {
	suite.Suite
}

func (s *p0970TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, powerfulIntegers(v.arg1, v.arg2, v.arg3))
	}
}

func TestP0970TestSuite(t *testing.T) {
	s := &p0970TestSuite{}
	suite.Run(t, s)
}
