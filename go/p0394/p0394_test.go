// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0394

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	target string
}

var values = []result{
	// {
	// 	arg1:   "3[a]2[bc]",
	// 	target: "aaabcbc",
	// },
	{
		arg1:   "3[a2[c]]",
		target: "accaccacc",
	},
	{
		arg1:   "2[abc]3[cd]ef",
		target: "abcabccdcdcdef",
	},
}

type p0394TestSuite struct {
	suite.Suite
}

func (s *p0394TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, decodeString(v.arg1))
	}
}

func TestP0394TestSuite(t *testing.T) {
	s := &p0394TestSuite{}
	suite.Run(t, s)
}
