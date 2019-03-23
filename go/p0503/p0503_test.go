// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0503

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
		arg1:   []int{1, 2, 1},
		target: []int{2, -1, 2},
	},
	{
		arg1:   []int{3, -2, -1},
		target: []int{-1, -1, 3},
	},
}

type p0503TestSuite struct {
	suite.Suite
}

func (s *p0503TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, nextGreaterElements(v.arg1))
	}
}

func TestP0503TestSuite(t *testing.T) {
	s := &p0503TestSuite{}
	suite.Run(t, s)
}
