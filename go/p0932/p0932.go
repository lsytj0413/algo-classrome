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

package p0932

import (
	"reflect"
	"unsafe"
)

// TODO: 需要证明
func beautifulArray(N int) []int {
	r := make([]int, 0, N)
	r = append(r, 1)
	for len(r) < N {
		tmp := make([]int, len(r))
		copy(tmp, r)

		h := (*reflect.SliceHeader)(unsafe.Pointer(&r))
		h.Len = 0
		for _, i := range tmp {
			r = append(r, 2*i-1)
		}
		for _, i := range tmp {
			r = append(r, 2*i)
		}
	}

	res := make([]int, 0, N)
	for _, v := range r {
		if v <= N {
			res = append(res, v)
		}
	}
	return res
}
