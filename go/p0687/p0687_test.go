// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0687

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
		arg1:   []int{5, 4, 5, 1, 1, 0xFFFFFFFF, 5},
		target: 2,
	},
	{
		arg1:   []int{5, 4, 5, 4, 4, 0xFFFFFFFF, 5},
		target: 2,
	},
	{
		arg1:   []int{4, -7, -3, 0xFFFFFFFF, 0xFFFFFFFF, -9, -3, 9, -7, -4, 0xFFFFFFFF, 6, 0xFFFFFFFF, -6, -6, 0xFFFFFFFF, 0xFFFFFFFF, 0, 6, 5, 0xFFFFFFFF, 9, 0xFFFFFFFF, 0xFFFFFFFF, -1, -4, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, -2},
		target: 1,
	},
}

type p0687TestSuite struct {
	suite.Suite
}

func (s *p0687TestSuite) Test() {
	for _, v := range values {
		t := comm.Trees(v.arg1)
		s.Equal(v.target, longestUnivaluePath(t))
	}
}

func TestP0687TestSuite(t *testing.T) {
	s := &p0687TestSuite{}
	suite.Run(t, s)
}
