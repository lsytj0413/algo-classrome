// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0395

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   int
	target int
}

var values = []result{
	{
		arg1:   "aaabb",
		arg2:   3,
		target: 3,
	},
	{
		arg1:   "ababbc",
		arg2:   2,
		target: 5,
	},
}

type p0395TestSuite struct {
	suite.Suite
}

func (s *p0395TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, longestSubstring(v.arg1, v.arg2))
	}
}

func TestP0395TestSuite(t *testing.T) {
	s := &p0395TestSuite{}
	suite.Run(t, s)
}
