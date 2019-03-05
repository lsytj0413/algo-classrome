// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0491

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []int
	target [][]int
}

var values = []result{
	{
		arg1: []int{4, 6, 7, 7},
		target: [][]int{
			[]int{4, 6},
			[]int{4, 6, 7},
			[]int{4, 7},
			[]int{6, 7},
			[]int{4, 6, 7, 7},
			[]int{4, 7, 7},
			[]int{6, 7, 7},
			[]int{7, 7},
		},
	},
}

type p0491TestSuite struct {
	suite.Suite
}

func (s *p0491TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findSubsequences(v.arg1))
	}
}

func TestP0491TestSuite(t *testing.T) {
	s := &p0491TestSuite{}
	suite.Run(t, s)
}
