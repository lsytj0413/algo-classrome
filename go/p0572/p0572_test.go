// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0572

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"algo/comm"
)

type result struct {
	arg1   []int
	arg2   []int
	target bool
}

var values = []result{
	{
		arg1:   []int{3, 4, 5, 1, 2},
		arg2:   []int{4, 1, 2},
		target: true,
	},
	{
		arg1:   []int{3, 4, 5, 1, 2, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0},
		arg2:   []int{4, 1, 2},
		target: false,
	},
	{
		arg1:   []int{3, 4, 5, 1, 0xFFFFFFFF, 2},
		arg2:   []int{3, 1, 2},
		target: false,
	},
	{
		arg1:   []int{3, 4, 5, 1, 2},
		arg2:   []int{4, 1, 2, 1},
		target: false,
	},
}

type p0572TestSuite struct {
	suite.Suite
}

func (s *p0572TestSuite) Test() {
	for _, v := range values {
		r := isSubtree(comm.Trees(v.arg1), comm.Trees(v.arg2))
		s.Equal(v.target, r)
	}
}

func TestP0572TestSuite(t *testing.T) {
	s := &p0572TestSuite{}
	suite.Run(t, s)
}
