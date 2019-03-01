// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0442

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
		arg1:   []int{4, 3, 2, 7, 8, 2, 3, 1},
		target: []int{2, 3},
	},
	{
		arg1:   []int{10, 2, 5, 10, 9, 1, 1, 4, 3, 7},
		target: []int{1, 10},
	},
	{
		arg1:   []int{9, 9, 4, 10, 8, 5, 2, 2, 7, 7},
		target: []int{2, 7, 9},
	},
}

type p0442TestSuite struct {
	suite.Suite
}

func (s *p0442TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findDuplicates(v.arg1))
	}
}

func TestP0442TestSuite(t *testing.T) {
	s := &p0442TestSuite{}
	suite.Run(t, s)
}
