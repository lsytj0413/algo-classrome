// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0671

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
		arg1:   []int{2, 2, 5, 0xFFFFFFFF, 0xFFFFFFFF, 5, 7},
		target: 5,
	},
	{
		arg1:   []int{2, 2, 2},
		target: -1,
	},
}

type p0671TestSuite struct {
	suite.Suite
}

func (s *p0671TestSuite) Test() {
	for _, v := range values {
		r := findSecondMinimumValue(comm.Trees(v.arg1))
		s.Equal(v.target, r)
	}
}

func TestP0671TestSuite(t *testing.T) {
	s := &p0671TestSuite{}
	suite.Run(t, s)
}
