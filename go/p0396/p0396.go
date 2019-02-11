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

package p0396

func maxRotateFunction(A []int) int {
	f0, sum := 0, 0
	for i := 0; i < len(A); i++ {
		f0 += A[i] * i
		sum += A[i]
	}

	max, fi := f0, f0
	for i := 1; i < len(A); i++ {
		fi += sum
		fi -= len(A) * A[len(A)-i]
		if fi > max {
			max = fi
		}
	}

	return max
}
