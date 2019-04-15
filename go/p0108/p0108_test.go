// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0108

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
		arg1:   []int{-10, -3, 0, 5, 9},
		target: []int{0, -3, 9, -10, 0xFFFFFFFF, 5},
	},
}

type p0108TestSuite struct {
	suite.Suite
}

func (s *p0108TestSuite) Test() {
	for _, v := range values {
		comm.AssertTreeEqual(s.Suite, comm.Trees(v.target), sortedArrayToBST(v.arg1))
	}
}

func TestP0108TestSuite(t *testing.T) {
	s := &p0108TestSuite{}
	suite.Run(t, s)
}
