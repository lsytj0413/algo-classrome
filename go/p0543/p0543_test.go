// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0543

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
		arg1:   []int{1, 2, 3, 4, 5},
		target: 3,
	},
}

type p0543TestSuite struct {
	suite.Suite
}

func (s *p0543TestSuite) Test() {
	for _, v := range values {
		r := diameterOfBinaryTree(comm.Trees(v.arg1))
		s.Equal(v.target, r)
	}
}

func TestP0543TestSuite(t *testing.T) {
	s := &p0543TestSuite{}
	suite.Run(t, s)
}
