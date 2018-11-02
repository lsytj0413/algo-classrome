// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0322

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	arg2   int
	target int
}

var values = []result{
	{
		arg1:   []int{1, 2, 5},
		arg2:   11,
		target: 3,
	},
	{
		arg1:   []int{2},
		arg2:   3,
		target: -1,
	},
	{
		arg1:   []int{2, 5, 10, 1},
		arg2:   27,
		target: 4,
	},
	// {
	// 	arg1:   []int{186, 419, 83, 408},
	// 	arg2:   6249,
	// 	target: 20,
	// },
	{
		arg1:   []int{1, 3, 5, 6},
		arg2:   11,
		target: 2,
	},
}

type p0322TestSuite struct {
	suite.Suite
}

func (s *p0322TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, coinChange(v.arg1, v.arg2))
	}
}

func TestP0322TestSuite(t *testing.T) {
	s := &p0322TestSuite{}
	suite.Run(t, s)
}
