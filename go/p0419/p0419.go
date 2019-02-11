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

package p0419

func countBattleships(board [][]byte) int {
	count := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			switch {
			case i == 0 && j == 0:
				if board[i][j] == 'X' {
					count++
				}
			case i == 0:
				if board[i][j] == 'X' && board[i][j-1] == '.' {
					count++
				}
			case j == 0:
				if board[i][j] == 'X' && board[i-1][j] == '.' {
					count++
				}
			default:
				if board[i][j] == 'X' && board[i-1][j] == '.' && board[i][j-1] == '.' {
					count++
				}
			}
		}
	}
	return count
}
