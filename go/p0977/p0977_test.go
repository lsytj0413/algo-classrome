// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0977

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	target []int
}

var values = []result{
	{
		arg1:   []int{-4, -1, 0, 3, 10},
		target: []int{0, 1, 9, 16, 100},
	},
	{
		arg1:   []int{-7, -3, 2, 3, 11},
		target: []int{4, 9, 9, 49, 121},
	},
}

type p0977TestSuite struct {
	suite.Suite
}

func (s *p0977TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, sortedSquares(v.arg1))
	}
}

func TestP0977TestSuite(t *testing.T) {
	s := &p0977TestSuite{}
	suite.Run(t, s)
}
