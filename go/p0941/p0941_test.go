// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0941

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
		arg1:   []int{2, 1},
		target: false,
	},
	{
		arg1:   []int{3, 5, 5},
		target: false,
	},
	{
		arg1:   []int{0, 3, 2, 1},
		target: true,
	},
}

type p0941TestSuite struct {
	suite.Suite
}

func (s *p0941TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, validMountainArray(v.arg1))
	}
}

func TestP0941TestSuite(t *testing.T) {
	s := &p0941TestSuite{}
	suite.Run(t, s)
}
