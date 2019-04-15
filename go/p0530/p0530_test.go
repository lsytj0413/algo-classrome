// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0530

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
		arg1:   []int{1, 0xFFFFFFFF, 3, 0xFFFFFFFF, 0xFFFFFFFF, 2},
		target: 1,
	},
}

type p0530TestSuite struct {
	suite.Suite
}

func (s *p0530TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, getMinimumDifference(comm.Trees(v.arg1)))
	}
}

func TestP0530TestSuite(t *testing.T) {
	s := &p0530TestSuite{}
	suite.Run(t, s)
}
