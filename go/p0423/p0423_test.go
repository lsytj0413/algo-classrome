// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0423

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	target string
}

var values = []result{
	{
		arg1:   "owoztneoer",
		target: "012",
	},
	{
		arg1:   "fviefuro",
		target: "45",
	},
	{
		arg1:   "seven",
		target: "7",
	},
	{
		arg1:   "zeroonetwothreefourfivesixseveneightnine",
		target: "0123456789",
	},
}

type p0423TestSuite struct {
	suite.Suite
}

func (s *p0423TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, originalDigits(v.arg1))
	}
}

func TestP0423TestSuite(t *testing.T) {
	s := &p0423TestSuite{}
	suite.Run(t, s)
}
