// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0953

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []string
	arg2   string
	target bool
}

var values = []result{
	{
		arg1:   []string{"hello", "leetcode"},
		arg2:   "hlabcdefgijkmnopqrstuvwxyz",
		target: true,
	},
	{
		arg1:   []string{"word", "world", "row"},
		arg2:   "worldabcefghijkmnpqstuvxyz",
		target: false,
	},
	{
		arg1:   []string{"apple", "app"},
		arg2:   "abcdefghijklmnopqrstuvwxyz",
		target: false,
	},
}

type p0953TestSuite struct {
	suite.Suite
}

func (s *p0953TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, isAlienSorted(v.arg1, v.arg2))
	}
}

func TestP0953TestSuite(t *testing.T) {
	s := &p0953TestSuite{}
	suite.Run(t, s)
}
