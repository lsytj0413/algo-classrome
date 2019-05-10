// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0989

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   int
	target []int
}

var values = []result{
	{
		arg1:   []int{1, 2, 0, 0},
		arg2:   34,
		target: []int{1, 2, 3, 4},
	},
	{
		arg1:   []int{2, 7, 4},
		arg2:   181,
		target: []int{4, 5, 5},
	},
	{
		arg1:   []int{2, 1, 5},
		arg2:   806,
		target: []int{1, 0, 2, 1},
	},
	{
		arg1:   []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		arg2:   1,
		target: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		arg1:   []int{0},
		arg2:   10000,
		target: []int{1, 0, 0, 0, 0},
	},
}

type p0989TestSuite struct {
	suite.Suite
}

func (s *p0989TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, addToArrayForm(v.arg1, v.arg2))
	}
}

func TestP0989TestSuite(t *testing.T) {
	s := &p0989TestSuite{}
	suite.Run(t, s)
}
