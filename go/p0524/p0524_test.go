// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0524

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   []string
	target string
}

var values = []result{
	{
		arg1:   "abpcplea",
		arg2:   []string{"ale", "apple", "monkey", "plea"},
		target: "apple",
	},
	{
		arg1:   "abpcplea",
		arg2:   []string{"a", "b", "c"},
		target: "a",
	},
}

type p0524TestSuite struct {
	suite.Suite
}

func (s *p0524TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findLongestWord(v.arg1, v.arg2))
	}
}

func TestP0524TestSuite(t *testing.T) {
	s := &p0524TestSuite{}
	suite.Run(t, s)
}
