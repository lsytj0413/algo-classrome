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

package comm

// Interval is a struct with start and end point
type Interval struct {
	Start int
	End   int
}

// Intervals for sort
type Intervals []Interval

// Len for sort.Interface
func (p Intervals) Len() int {
	return len(p)
}

// Swap for sort.Interface
func (p Intervals) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Less for sort.Interface
func (p Intervals) Less(i, j int) bool {
	return p[i].Start < p[j].Start
}
