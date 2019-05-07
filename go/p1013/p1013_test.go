// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p1013

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	target bool
}

var values = []result{
	{
		arg1:   []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1},
		target: true,
	},
	{
		arg1:   []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1},
		target: false,
	},
	{
		arg1:   []int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4},
		target: true,
	},
	{
		arg1:   []int{18, 12, -18, 18, -19, -1, 10, 10},
		target: true,
	},
}

type p1013TestSuite struct {
	suite.Suite
}

func (s *p1013TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, canThreePartsEqualSum(v.arg1))
	}
}

func TestP1013TestSuite(t *testing.T) {
	s := &p1013TestSuite{}
	suite.Run(t, s)
}
