// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0385

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1   string
	target int
}

var values = []result{
	{
		arg1:   "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext",
		target: 32,
	},
	{
		arg1:   "dir\n\tfile.txt",
		target: 12,
	},
	{
		arg1:   "dir\n    file.txt",
		target: 12,
	},
	{
		arg1:   "dir\n        file.txt",
		target: 16,
	},
	{
		arg1:   "a\n\tb.txt\na2\n\tb2.txt",
		target: 9,
	},
	{
		arg1:   "a\n\tb\n\t\tc.txt\n\taaaa.txt",
		target: 10,
	},
}

type p0385TestSuite struct {
	suite.Suite
}

func (s *p0385TestSuite) Test() {
	// for _, v := range values {
	// 	s.Equal(v.target, lengthLongestPath(v.arg1))
	// }
	deserialize("[-1,-2]")
}

func TestP0385TestSuite(t *testing.T) {
	s := &p0385TestSuite{}
	suite.Run(t, s)
}
