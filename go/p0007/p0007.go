// Copyright (c) 2018 soren yang
//
// Licensed under the MIT License
// you may not use this file except in complicance with the License.
// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0007

import (
	"math"
)

func reverse(x int) (num int) {
	for x != 0 {
		num = num*10 + x%10
		x = x / 10
	}
	// 使用 math 包中定义好的最大最小值
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}

	return
}
