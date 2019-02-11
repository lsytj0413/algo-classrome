// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0421

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
		arg1:   []int{3, 10, 5, 25, 2, 8},
		target: 28,
	},
}

type p0421TestSuite struct {
	suite.Suite
}

func (s *p0421TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findMaximumXOR(v.arg1))
	}
}

func TestP0421TestSuite(t *testing.T) {
	s := &p0421TestSuite{}
	suite.Run(t, s)
}
