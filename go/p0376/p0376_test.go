// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0376

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	target int
}

var values = []result{
	{
		arg1:   []int{1, 7, 4, 9, 2, 5},
		target: 6,
	},
	{
		arg1:   []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8},
		target: 7,
	},
	{
		arg1:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		target: 2,
	},
}

type p0376TestSuite struct {
	suite.Suite
}

func (s *p0376TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, wiggleMaxLength(v.arg1))
	}
}

func TestP0376TestSuite(t *testing.T) {
	s := &p0376TestSuite{}
	suite.Run(t, s)
}
