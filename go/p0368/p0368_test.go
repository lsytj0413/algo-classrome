// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0368

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	target []int
}

var values = []result{
	{
		arg1:   []int{1, 2, 3},
		target: []int{1, 2},
	},
	{
		arg1:   []int{1, 2, 4, 8},
		target: []int{1, 2, 4, 8},
	},
	{
		arg1:   []int{4, 8, 10, 240},
		target: []int{4, 8, 240},
	},
	{
		arg1:   []int{},
		target: []int{},
	},
}

type p0368TestSuite struct {
	suite.Suite
}

func (s *p0368TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, largestDivisibleSubset(v.arg1))
	}
}

func TestP0368TestSuite(t *testing.T) {
	s := &p0368TestSuite{}
	suite.Run(t, s)
}
