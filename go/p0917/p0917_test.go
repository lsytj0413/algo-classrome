// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0917

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
		arg1:   "ab-cd",
		target: "dc-ba",
	},
	{
		arg1:   "a-bC-dEf-ghIj",
		target: "j-Ih-gfE-dCba",
	},
	{
		arg1:   "Test1ng-Leet=code-Q!",
		target: "Qedo1ct-eeLg=ntse-T!",
	},
}

type p0917TestSuite struct {
	suite.Suite
}

func (s *p0917TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, reverseOnlyLetters(v.arg1))
	}
}

func TestP0917TestSuite(t *testing.T) {
	s := &p0917TestSuite{}
	suite.Run(t, s)
}
