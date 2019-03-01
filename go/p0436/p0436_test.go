// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0436

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []Interval
	target []int
}

var values = []result{
	{
		arg1: []Interval{
			{Start: 1, End: 2},
		},
		target: []int{-1},
	},
	{
		arg1: []Interval{
			{Start: 3, End: 4},
			{Start: 2, End: 3},
			{Start: 1, End: 2},
		},
		target: []int{-1, 0, 1},
	},
	{
		arg1: []Interval{
			{Start: 1, End: 4},
			{Start: 2, End: 3},
			{Start: 3, End: 4},
		},
		target: []int{-1, 2, -1},
	},
}

type p0436TestSuite struct {
	suite.Suite
}

func (s *p0436TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findRightInterval(v.arg1))
	}
}

func TestP0436TestSuite(t *testing.T) {
	s := &p0436TestSuite{}
	suite.Run(t, s)
}
