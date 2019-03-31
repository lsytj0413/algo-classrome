// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0523

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   int
	target bool
}

var values = []result{
	{
		arg1:   []int{23, 2, 4, 6, 7},
		arg2:   6,
		target: true,
	},
	{
		arg1:   []int{23, 2, 6, 4, 7},
		arg2:   6,
		target: true,
	},
	{
		arg1:   []int{23, 2, 6, 4, 7},
		arg2:   0,
		target: false,
	},
	{
		arg1:   []int{23, 2, 4, 6, 7},
		arg2:   -6,
		target: true,
	},
}

type p0523TestSuite struct {
	suite.Suite
}

func (s *p0523TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, checkSubarraySum(v.arg1, v.arg2))
	}
}

func TestP0523TestSuite(t *testing.T) {
	s := &p0523TestSuite{}
	suite.Run(t, s)
}
