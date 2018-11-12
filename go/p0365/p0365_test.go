// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0365

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	arg2   int
	arg3   int
	target bool
}

var values = []result{
	{
		arg1:   3,
		arg2:   5,
		arg3:   4,
		target: true,
	},
	{
		arg1:   2,
		arg2:   6,
		arg3:   5,
		target: false,
	},
}

type p0365TestSuite struct {
	suite.Suite
}

func (s *p0365TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, canMeasureWater(v.arg1, v.arg2, v.arg3))
	}
}

func TestP0365TestSuite(t *testing.T) {
	s := &p0365TestSuite{}
	suite.Run(t, s)
}
