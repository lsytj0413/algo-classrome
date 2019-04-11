// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0748

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	arg2   []string
	target string
}

var values = []result{
	{
		arg1:   "1s3 PSt",
		arg2:   []string{"step", "steps", "stripe", "stepple"},
		target: "steps",
	},
	{
		arg1:   "1s3 456",
		arg2:   []string{"looks", "pest", "stew", "show"},
		target: "pest",
	},
}

type p0748TestSuite struct {
	suite.Suite
}

func (s *p0748TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, shortestCompletingWord(v.arg1, v.arg2))
	}
}

func TestP0748TestSuite(t *testing.T) {
	s := &p0748TestSuite{}
	suite.Run(t, s)
}
