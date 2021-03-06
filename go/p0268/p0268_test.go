// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0268

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
		arg1:   []int{3, 0, 1},
		target: 2,
	},
	{
		arg1:   []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
		target: 8,
	},
}

type p0268TestSuite struct {
	suite.Suite
}

func (s *p0268TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, missingNumber(v.arg1))
	}
}

func TestP0268TestSuite(t *testing.T) {
	s := &p0268TestSuite{}
	suite.Run(t, s)
}
