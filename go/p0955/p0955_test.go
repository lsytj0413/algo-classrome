// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0955

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   []string
	target int
}

var values = []result{
	{
		arg1:   []string{"ca", "bb", "ac"},
		target: 1,
	},
	{
		arg1:   []string{"xc", "yb", "za"},
		target: 0,
	},
	{
		arg1:   []string{"zyx", "wvu", "tsr"},
		target: 3,
	},
	{
		arg1:   []string{"xga", "xfb", "yfa"},
		target: 1,
	},
	{
		arg1:   []string{"abx", "agz", "bgc", "bfc"},
		target: 1,
	},
	{
		arg1:   []string{"jsebodtwc", "cnneoytml", "eeneuyzlu", "ewpnmtivg"},
		target: 2,
	},
	{
		arg1: []string{
			"mwdagrdrrzglubcydvrkvedsrdogumgrymbnbjlfwpudlzeiqm", "oxlasuiyboyetbclkzlppkfjdrmhpwsuwfcmvbbgpqiokxgubw", "egyeqtqqqhkiqwdxbmkmydpfnkpdoltinogagyxxmwqoiijcwo", "lswlrvgrifuastgfxvlldpdbamsrwrojqgjxephchdqpllombu", "cjucfhqkowdbdfsdyumbnbuuhgiyrsysstjgdsmwqwhzkxppvx", "grccjwniaznlusvedsbbvihqcnyehjykbznibckqbxzrsbgbet", "uqhwondfirzqxpitvjcedphjxuybgiolgaphlcutrtmejedddx", "askalyaxvhflvqhalnnmhqimnxlpjqlxuawhreasnysilxpxuh", "qpdlpivobfpdyqgktfqnaheobpvibyijopwcrabxcqzkuifbso", "nkipclxiwvotmsacnscnzyvbgnvakigoczzqvkknkabisafksz"},
		target: 37,
	},
}

type p0955TestSuite struct {
	suite.Suite
}

func (s *p0955TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, minDeletionSize(v.arg1))
	}
}

func TestP0955TestSuite(t *testing.T) {
	s := &p0955TestSuite{}
	suite.Run(t, s)
}
