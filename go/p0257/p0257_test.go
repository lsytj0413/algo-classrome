// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0257

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"
	
	"algo/comm"
)

type result struct {
	arg1   []int
	target []string
}

var values = []result{
	{
		arg1:   []int{1, 2, 3, 0xFFFFFFFF, 5},
		target: []string{"1->2->5", "1->3"},
	},
}

type p0257TestSuite struct {
	suite.Suite
}

func (s *p0257TestSuite) Test() {
	for _, v := range values {
		r := binaryTreePaths(comm.Trees(v.arg1))
		sort.Strings(r)
		sort.Strings(v.target)
		s.Equal(v.target, r)
	}
}

func TestP0257TestSuite(t *testing.T) {
	s := &p0257TestSuite{}
	suite.Run(t, s)
}
