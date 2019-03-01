// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0451

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
		arg1:   "tree",
		target: "eert",
	},
	{
		arg1:   "cccaaa",
		target: "aaaccc",
	},
	{
		arg1:   "Aabb",
		target: "bbAa",
	},
}

type p0451TestSuite struct {
	suite.Suite
}

func (s *p0451TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, frequencySort(v.arg1))
	}
}

func TestP0451TestSuite(t *testing.T) {
	s := &p0451TestSuite{}
	suite.Run(t, s)
}
