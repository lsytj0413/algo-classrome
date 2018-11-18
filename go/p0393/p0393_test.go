// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0393

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
		arg1:   []int{197, 130, 1},
		target: true,
	},
	{
		arg1:   []int{235, 140, 4},
		target: false,
	},
	{
		arg1:   []int{255},
		target: false,
	},
	{
		arg1:   []int{248, 130, 130, 130},
		target: false,
	},
}

type p0393TestSuite struct {
	suite.Suite
}

func (s *p0393TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, validUtf8(v.arg1))
	}
}

func TestP0393TestSuite(t *testing.T) {
	s := &p0393TestSuite{}
	suite.Run(t, s)
}
