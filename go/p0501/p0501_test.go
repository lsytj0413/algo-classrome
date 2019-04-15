// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0501

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"algo/comm"
)

type result struct {
	arg1   []int
	target []int
}

var values = []result{
	{
		arg1:   []int{1, 0xFFFFFFFF, 2, 0xFFFFFFFF, 0xFFFFFFFF, 2},
		target: []int{2},
	},
	{
		arg1:   []int{2, 0xFFFFFFFF, 3, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 4, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 5, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 6},
		target: []int{2, 3, 4, 5, 6},
	},
}

type p0501TestSuite struct {
	suite.Suite
}

func (s *p0501TestSuite) Test() {
	for _, v := range values {
		r := findMode(comm.Trees(v.arg1))
		s.Equal(v.target, r)
	}
}

func TestP0501TestSuite(t *testing.T) {
	s := &p0501TestSuite{}
	suite.Run(t, s)
}
