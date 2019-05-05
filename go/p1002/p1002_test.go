// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p1002

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []string
	target []string
}

var values = []result{
	{
		arg1:   []string{"bella", "label", "roller"},
		target: []string{"e", "l", "l"},
	},
	{
		arg1:   []string{"cool", "lock", "cook"},
		target: []string{"c", "o"},
	},
}

type p1002TestSuite struct {
	suite.Suite
}

func (s *p1002TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, commonChars(v.arg1))
	}
}

func TestP1002TestSuite(t *testing.T) {
	s := &p1002TestSuite{}
	suite.Run(t, s)
}
