// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0324

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
		arg1:   []int{1, 5, 1, 1, 6, 4},
		target: []int{1, 6, 1, 5, 1, 4},
	},
	{
		arg1:   []int{1, 3, 2, 2, 3, 1},
		target: []int{2, 3, 1, 3, 1, 2},
	},
	{
		arg1:   []int{4, 5, 5, 6},
		target: []int{5, 6, 4, 5},
	},
}

type p0324TestSuite struct {
	suite.Suite
}

func (s *p0324TestSuite) Test() {
	for _, v := range values {
		wiggleSort(v.arg1)
		s.Equal(v.target, v.arg1)
	}
}

func TestP0324TestSuite(t *testing.T) {
	s := &p0324TestSuite{}
	suite.Run(t, s)
}
