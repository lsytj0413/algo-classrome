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

package p0498

func findDiagonalOrder(matrix [][]int) []int {
	row := len(matrix)
	if row == 0 {
		return nil
	}
	col := len(matrix[0])
	if col == 0 {
		return nil
	}

	r := make([]int, 0, row*col)
	cr, cc, i := 0, 0, 0
	for i < row*col {
		r = append(r, matrix[cr][cc])
		switch (cr + cc) & 0x01 {
		case 1:
			if cr < row-1 && cc > 0 {
				cr++
				cc--
			} else {
				if cr == row-1 && cc == 0 {
					cc++
				} else if cr == row-1 {
					cc++
				} else if cc == 0 {
					cr++
				}
			}
		case 0:
			if cr > 0 && cc < col-1 {
				cr--
				cc++
			} else {
				if cr == 0 && cc == col-1 {
					cr++
				} else if cr == 0 {
					cc++
				} else if cc == col-1 {
					cr++
				}
			}
		}
		i++
	}
	return r
}
