// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0961

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	target int
}

var values = []result{
	{
		arg1:   []int{1, 2, 3, 3},
		target: 3,
	},
	{
		arg1:   []int{2, 1, 2, 5, 3, 2},
		target: 2,
	},
	{
		arg1:   []int{5, 1, 5, 2, 5, 3, 5, 4},
		target: 5,
	},
	{
		arg1:   []int{9, 5, 6, 9},
		target: 9,
	},
}

type p0961TestSuite struct {
	suite.Suite
}

func (s *p0961TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, repeatedNTimes(v.arg1))
	}
}

func TestP0961TestSuite(t *testing.T) {
	s := &p0961TestSuite{}
	suite.Run(t, s)
}
