// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0538

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
		arg1:   []int{5, 2, 13},
		target: []int{18, 20, 13},
	},
	{
		arg1:   []int{5, 2, 13, 0, 3, 11, 14},
		target: []int{43, 48, 27, 48, 46, 38, 14},
	},
	{
		arg1:   []int{0, -1, 2, -3, 0xFFFFFFFF, 0xFFFFFFFF, 4},
		target: []int{6, 5, 6, 2, 0xFFFFFFFF, 0xFFFFFFFF, 4},
	},
}

type p0538TestSuite struct {
	suite.Suite
}

func (s *p0538TestSuite) Test() {
	for _, v := range values {
		r := convertBST(comm.Trees(v.arg1))
		comm.AssertTreeEqual(s.Suite, comm.Trees(v.target), r)
	}
}

func TestP0538TestSuite(t *testing.T) {
	s := &p0538TestSuite{}
	suite.Run(t, s)
}
