// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0944

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []string
	target int
}

var values = []result{
	{
		arg1:   []string{"cba", "daf", "ghi"},
		target: 1,
	},
	{
		arg1:   []string{"a", "b"},
		target: 0,
	},
	{
		arg1:   []string{"zyx", "wvu", "tsr"},
		target: 3,
	},
}

type p0944TestSuite struct {
	suite.Suite
}

func (s *p0944TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, minDeletionSize(v.arg1))
	}
}

func TestP0944TestSuite(t *testing.T) {
	s := &p0944TestSuite{}
	suite.Run(t, s)
}
