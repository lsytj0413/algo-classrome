// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0307

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	target bool
}

var values = []result{
	{
		arg1:   27,
		target: true,
	},
	{
		arg1:   0,
		target: false,
	},
	{
		arg1:   9,
		target: true,
	},
	{
		arg1:   45,
		target: false,
	},
}

type p0307TestSuite struct {
	suite.Suite
}

func (s *p0307TestSuite) Test() {
	// for _, v := range values {
	// 	s.Equal(v.target, isPowerOfThree(v.arg1))
	// }
	n := Constructor([]int{1, 3, 5})
	s.Equal(9, n.SumRange(0, 2))
	n.Update(1, 2)
	s.Equal(8, n.SumRange(0, 2))
}

func TestP0307TestSuite(t *testing.T) {
	s := &p0307TestSuite{}
	suite.Run(t, s)
}
