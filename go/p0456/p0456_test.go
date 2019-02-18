// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0456

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	target bool
}

var values = []result{
	{
		arg1:   []int{1, 2, 3, 4},
		target: false,
	},
	{
		arg1:   []int{3, 1, 4, 2},
		target: true,
	},
	{
		arg1:   []int{-1, 3, 2, 0},
		target: true,
	},
}

type p0456TestSuite struct {
	suite.Suite
}

func (s *p0456TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, find132pattern(v.arg1))
	}
}

func TestP0456TestSuite(t *testing.T) {
	s := &p0456TestSuite{}
	suite.Run(t, s)
}
