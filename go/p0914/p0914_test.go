// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0914

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
		arg1:   []int{1, 2, 3, 4, 4, 3, 2, 1},
		target: true,
	},
	{
		arg1:   []int{1, 1, 1, 2, 2, 2, 3, 3},
		target: false,
	},
	{
		arg1:   []int{1},
		target: false,
	},
	{
		arg1:   []int{1, 1},
		target: true,
	},
	{
		arg1:   []int{1, 1, 2, 2, 2, 2},
		target: true,
	},
	{
		arg1:   []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2},
		target: true,
	},
}

type p0914TestSuite struct {
	suite.Suite
}

func (s *p0914TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, hasGroupsSizeX(v.arg1))
	}
}

func TestP0914TestSuite(t *testing.T) {
	s := &p0914TestSuite{}
	suite.Run(t, s)
}
