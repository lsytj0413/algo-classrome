// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p1047

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	target string
}

var values = []result{
	{
		arg1:   "abbaca",
		target: "ca",
	},
	{
		arg1:   "aa",
		target: "",
	},
	{
		arg1:   "aabc",
		target: "bc",
	},
}

type p1047TestSuite struct {
	suite.Suite
}

func (s *p1047TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, removeDuplicates(v.arg1))
	}
}

func TestP1047TestSuite(t *testing.T) {
	s := &p1047TestSuite{}
	suite.Run(t, s)
}
