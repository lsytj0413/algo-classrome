// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p1014

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
		arg1:   []int{8, 1, 5, 2, 6},
		target: 11,
	},
}

type p1014TestSuite struct {
	suite.Suite
}

func (s *p1014TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, maxScoreSightseeingPair(v.arg1))
	}
}

func TestP1014TestSuite(t *testing.T) {
	s := &p1014TestSuite{}
	suite.Run(t, s)
}
