// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0735

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
		arg1:   []int{5, 10, -5},
		target: []int{5, 10},
	},
	{
		arg1:   []int{8, -8},
		target: []int{},
	},
	{
		arg1:   []int{10, 2, -5},
		target: []int{10},
	},
	{
		arg1:   []int{-2, -1, 1, 2},
		target: []int{-2, -1, 1, 2},
	},
}

type p0735TestSuite struct {
	suite.Suite
}

func (s *p0735TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, asteroidCollision(v.arg1))
	}
}

func TestP0735TestSuite(t *testing.T) {
	s := &p0735TestSuite{}
	suite.Run(t, s)
}
