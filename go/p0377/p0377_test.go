// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0377

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   int
	target int
}

var values = []result{
	{
		arg1:   []int{1, 2, 3},
		arg2:   4,
		target: 7,
	},
}

type p0377TestSuite struct {
	suite.Suite
}

func (s *p0377TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, combinationSum4(v.arg1, v.arg2))
	}
}

func TestP0377TestSuite(t *testing.T) {
	s := &p0377TestSuite{}
	suite.Run(t, s)
}
