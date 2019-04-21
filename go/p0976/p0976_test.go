// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0976

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
		arg1:   []int{2, 1, 2},
		target: 5,
	},
	{
		arg1:   []int{1, 2, 1},
		target: 0,
	},
	{
		arg1:   []int{3, 2, 3, 4},
		target: 10,
	},
	{
		arg1:   []int{3, 6, 2, 3},
		target: 8,
	},
}

type p0976TestSuite struct {
	suite.Suite
}

func (s *p0976TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, largestPerimeter(v.arg1))
	}
}

func TestP0976TestSuite(t *testing.T) {
	s := &p0976TestSuite{}
	suite.Run(t, s)
}
