// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0932

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	target []int
}

var values = []result{
	{
		arg1:   4,
		target: []int{1, 3, 2, 4},
	},
	{
		arg1:   5,
		target: []int{1, 5, 3, 2, 4},
	},
}

type p0932TestSuite struct {
	suite.Suite
}

func (s *p0932TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, beautifulArray(v.arg1))
	}
}

func TestP0932TestSuite(t *testing.T) {
	s := &p0932TestSuite{}
	suite.Run(t, s)
}
