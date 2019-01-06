// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0402

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   int
	target string
}

var values = []result{
	{
		arg1:   "1432219",
		arg2:   3,
		target: "1219",
	},
	{
		arg1:   "10200",
		arg2:   1,
		target: "200",
	},
	{
		arg1:   "9",
		arg2:   1,
		target: "0",
	},
}

type p0402TestSuite struct {
	suite.Suite
}

func (s *p0402TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, removeKdigits(v.arg1, v.arg2))
	}
}

func TestP0402TestSuite(t *testing.T) {
	s := &p0402TestSuite{}
	suite.Run(t, s)
}
