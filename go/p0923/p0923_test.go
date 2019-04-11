// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0923

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
		arg1:   []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
		arg2:   8,
		target: 20,
	},
	{
		arg1:   []int{1, 1, 2, 2, 2, 2},
		arg2:   5,
		target: 12,
	},
}

type p0923TestSuite struct {
	suite.Suite
}

func (s *p0923TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, threeSumMulti(v.arg1, v.arg2))
	}
}

func TestP0923TestSuite(t *testing.T) {
	s := &p0923TestSuite{}
	suite.Run(t, s)
}
