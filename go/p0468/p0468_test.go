// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0468

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
		arg1:   "172.16.254.1",
		target: "IPv4",
	},
	{
		arg1:   "2001:0db8:85a3:0:0:8A2E:0370:7334",
		target: "IPv6",
	},
	{
		arg1:   "256.256.256.256",
		target: "Neither",
	},
	{
		arg1:   "01.01.01.01",
		target: "Neither",
	},
	{
		arg1:   "0.0.0.-0",
		target: "Neither",
	},
	{
		arg1:   "2001:0db8:85a3:00000:0:8A2E:0370:7334",
		target: "Neither",
	},
	{
		arg1:   "2001:0db8:85a3:0000:0:8A2E:0370:733a",
		target: "IPv6",
	},
	{
		arg1:   "1081:db8:85a3:01:-0:8A2E:0370:7334",
		target: "Neither",
	},
	{
		arg1:   "2001:0db8:85a3::8A2E:0370:7334",
		target: "Neither",
	},
	{
		arg1:   "02001:0db8:85a3:0000:0000:8a2e:0370:7334",
		target: "Neither",
	},
}

type p0468TestSuite struct {
	suite.Suite
}

func (s *p0468TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, validIPAddress(v.arg1))
	}
}

func TestP0468TestSuite(t *testing.T) {
	s := &p0468TestSuite{}
	suite.Run(t, s)
}
