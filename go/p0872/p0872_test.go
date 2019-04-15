// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0872

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
		arg1:   []int{3, 5, 1, 6, 2, 9, 8, 0xFFFFFFFF, 0xFFFFFFFF, 7, 4},
		arg2:   []int{3, 5, 1, 6, 7, 4, 2, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 9, 8},
		target: true,
	},
}

type p0872TestSuite struct {
	suite.Suite
}

func (s *p0872TestSuite) Test() {
	for _, v := range values {
		r := leafSimilar(comm.Trees(v.arg1), comm.Trees(v.arg2))
		s.Equal(v.target, r)
	}
}

func TestP0872TestSuite(t *testing.T) {
	s := &p0872TestSuite{}
	suite.Run(t, s)
}
