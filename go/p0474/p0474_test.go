// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0474

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []string
	arg2   int
	arg3   int
	target int
}

var values = []result{
	{
		arg1:   []string{"10", "0001", "111001", "1", "0"},
		arg2:   5,
		arg3:   3,
		target: 4,
	},
	{
		arg1:   []string{"10", "0", "1"},
		arg2:   1,
		arg3:   1,
		target: 2,
	},
}

type p0474TestSuite struct {
	suite.Suite
}

func (s *p0474TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findMaxForm(v.arg1, v.arg2, v.arg3))
	}
}

func TestP0474TestSuite(t *testing.T) {
	s := &p0474TestSuite{}
	suite.Run(t, s)
}
