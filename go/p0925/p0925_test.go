// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0925

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   string
	target bool
}

var values = []result{
	{
		arg1:   "alex",
		arg2:   "aaleex",
		target: true,
	},
	{
		arg1:   "saeed",
		arg2:   "ssaaedd",
		target: false,
	},
	{
		arg1:   "leelee",
		arg2:   "lleeelee",
		target: true,
	},
	{
		arg1:   "laiden",
		arg2:   "laiden",
		target: true,
	},
}

type p0925TestSuite struct {
	suite.Suite
}

func (s *p0925TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, isLongPressedName(v.arg1, v.arg2))
	}
}

func TestP0925TestSuite(t *testing.T) {
	s := &p0925TestSuite{}
	suite.Run(t, s)
}
