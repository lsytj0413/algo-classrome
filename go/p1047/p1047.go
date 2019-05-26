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

package p1047

func removeDuplicates(S string) string {
	arr, w := []byte(S), 0

	for i := 0; i < len(arr); i++ {
		if w > 0 && arr[w-1] == arr[i] {
			w--
			continue
		}

		arr[w] = arr[i]
		w++
	}

	return string(arr[0:w])
}
