// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0467

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	target int
}

var values = []result{
	{
		arg1:   "a",
		target: 1,
	},
	{
		arg1:   "cac",
		target: 2,
	},
	{
		arg1:   "zab",
		target: 6,
	},
}

type p0467TestSuite struct {
	suite.Suite
}

func (s *p0467TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findSubstringInWraproundString(v.arg1))
	}
}

func TestP0467TestSuite(t *testing.T) {
	s := &p0467TestSuite{}
	suite.Run(t, s)
}
