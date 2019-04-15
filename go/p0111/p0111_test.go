// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0111

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"algo/comm"
)

type result struct {
	arg1   []int
	target int
}

var values = []result{
	{
		arg1:   []int{3, 9, 20, 0xFFFFFFFF, 0xFFFFFFFF, 15, 7},
		target: 2,
	},
	{
		arg1:   []int{1, 2},
		target: 2,
	},
	{
		arg1:   []int{1, 2, 3, 4, 5},
		target: 2,
	},
}

type p0111TestSuite struct {
	suite.Suite
}

func (s *p0111TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, minDepth(comm.Trees(v.arg1)))
	}
}

func TestP0111TestSuite(t *testing.T) {
	s := &p0111TestSuite{}
	suite.Run(t, s)
}
