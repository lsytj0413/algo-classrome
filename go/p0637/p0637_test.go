// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0637

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"algo/comm"
)

type result struct {
	arg1   []int
	target []float64
}

var values = []result{
	{
		arg1:   []int{3, 9, 20, 0xFFFFFFFF, 0xFFFFFFFF, 15, 7},
		target: []float64{3, 14.5, 11},
	},
}

type p0637TestSuite struct {
	suite.Suite
}

func (s *p0637TestSuite) Test() {
	for _, v := range values {
		r := averageOfLevels(comm.Trees(v.arg1))
		s.Equal(v.target, r)
	}
}

func TestP0637TestSuite(t *testing.T) {
	s := &p0637TestSuite{}
	suite.Run(t, s)
}
