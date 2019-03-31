// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0526

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   int
	target int
}

var values = []result{
	{
		arg1:   1,
		target: 1,
	},
	{
		arg1:   2,
		target: 2,
	},
	{
		arg1:   3,
		target: 3,
	},
	{
		arg1:   4,
		target: 8,
	},
	{
		arg1:   5,
		target: 10,
	},
	{
		arg1:   6,
		target: 36,
	},
	{
		arg1:   7,
		target: 41,
	},
	{
		arg1:   8,
		target: 132,
	},
	{
		arg1:   9,
		target: 250,
	},
	{
		arg1:   10,
		target: 700,
	},
	{
		arg1:   11,
		target: 750,
	},
	{
		arg1:   12,
		target: 4010,
	},
	{
		arg1:   13,
		target: 4237,
	},
	{
		arg1:   14,
		target: 10680,
	},
	{
		arg1:   15,
		target: 24679,
	},
}

type p0526TestSuite struct {
	suite.Suite
}

func (s *p0526TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, countArrangement(v.arg1))
	}
}

func TestP0526TestSuite(t *testing.T) {
	s := &p0526TestSuite{}
	suite.Run(t, s)
}
