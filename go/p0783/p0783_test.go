// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0783

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
	// {
	// 	arg1:   []int{4, 2, 6, 1, 3},
	// 	target: 1,
	// },
	{
		arg1:   []int{1, 0, 48, 0xFFFFFFFF, 0xFFFFFFFF, 12, 49},
		target: 1,
	},
}

type p0783TestSuite struct {
	suite.Suite
}

func (s *p0783TestSuite) Test() {
	for _, v := range values {
		r := minDiffInBST(comm.Trees(v.arg1))
		s.Equal(v.target, r)
	}
}

func TestP0783TestSuite(t *testing.T) {
	s := &p0783TestSuite{}
	suite.Run(t, s)
}
