// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p1003

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	target bool
}

var values = []result{
	{
		arg1:   "aabcbc",
		target: true,
	},
	{
		arg1:   "abcabcababcc",
		target: true,
	},
	{
		arg1:   "abccba",
		target: false,
	},
	{
		arg1:   "cababc",
		target: false,
	},
	{
		arg1:   "abcd",
		target: false,
	},
	{
		arg1:   "aba",
		target: false,
	},
	{
		arg1:   "aabcbcabc",
		target: true,
	},
	{
		arg1:   "aaabcbcbc",
		target: true,
	},
}

type p1003TestSuite struct {
	suite.Suite
}

func (s *p1003TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, isValid(v.arg1))
	}
}

func TestP1003TestSuite(t *testing.T) {
	s := &p1003TestSuite{}
	suite.Run(t, s)
}
