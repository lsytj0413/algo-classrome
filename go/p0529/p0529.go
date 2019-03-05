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

package p0529

import "container/list"

func updateBoard(board [][]byte, click []int) [][]byte {
	if click[0] < 0 || click[0] >= len(board) {
		return board
	}
	if click[1] < 0 || click[1] >= len(board[0]) {
		return board
	}

	row, col := click[0], click[1]
	v := board[row][col]
	switch v {
	case 'M':
		board[row][col] = 'X'
		return board
	}

	type Pos struct {
		row int
		col int
	}

	l := list.New()
	l.PushBack(Pos{
		row: row,
		col: col,
	})
	for l.Len() != 0 {
		e := l.Front()
		pos := e.Value.(Pos)

		v = board[pos.row][pos.col]
		switch v {
		case 'E':
			count := 0
			poses := make([]Pos, 0, 8)

			// first, the curren row
			if pos.col > 0 {
				if board[pos.row][pos.col-1] == 'M' {
					count++
				} else if board[pos.row][pos.col-1] == 'E' {
					poses = append(poses, Pos{row: pos.row, col: pos.col - 1})
				}
			}
			if pos.col < len(board[0])-1 {
				if board[pos.row][pos.col+1] == 'M' {
					count++
				} else if board[pos.row][pos.col+1] == 'E' {
					poses = append(poses, Pos{row: pos.row, col: pos.col + 1})
				}
			}

			// second, the upper row
			if pos.row > 0 {
				if pos.col > 0 {

					if board[pos.row-1][pos.col-1] == 'M' {
						count++
					} else if board[pos.row-1][pos.col-1] == 'E' {
						poses = append(poses, Pos{row: pos.row - 1, col: pos.col - 1})
					}
				}
				if pos.col < len(board[0])-1 {
					if board[pos.row-1][pos.col+1] == 'M' {
						count++
					} else if board[pos.row-1][pos.col+1] == 'E' {
						poses = append(poses, Pos{row: pos.row - 1, col: pos.col + 1})
					}
				}

				if board[pos.row-1][pos.col] == 'M' {
					count++
				} else if board[pos.row-1][pos.col] == 'E' {
					poses = append(poses, Pos{row: pos.row - 1, col: pos.col})
				}
			}

			// final, the below row
			if pos.row < len(board)-1 {
				if pos.col > 0 {
					if board[pos.row+1][pos.col-1] == 'M' {
						count++
					} else if board[pos.row+1][pos.col-1] == 'E' {
						poses = append(poses, Pos{row: pos.row + 1, col: pos.col - 1})
					}
				}
				if pos.col < len(board[0])-1 {
					if board[pos.row+1][pos.col+1] == 'M' {
						count++
					} else if board[pos.row+1][pos.col+1] == 'E' {
						poses = append(poses, Pos{row: pos.row + 1, col: pos.col + 1})
					}
				}

				if board[pos.row+1][pos.col] == 'M' {
					count++
				} else if board[pos.row+1][pos.col] == 'E' {
					poses = append(poses, Pos{row: pos.row + 1, col: pos.col})
				}
			}

			if count == 0 {
				board[pos.row][pos.col] = 'B'

				for i := range poses {
					l.PushBack(poses[i])
				}
			} else {
				board[pos.row][pos.col] = '0' + byte(count)
			}
		}

		l.Remove(e)
	}

	return board
}
