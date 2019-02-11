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

package p0399

// TODO: 增加 bfs 解法
// TODO: 增加并查集解法
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string][]string, 0)
	weight := make(map[[2]string]float64, 0)

	for i, v := range equations {
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
		weight[[2]string{v[0], v[1]}] = values[i]
		weight[[2]string{v[1], v[0]}] = 1 / values[i]
	}

	var _dfs func(string, string, map[string]bool) float64
	_dfs = func(start string, end string, visited map[string]bool) float64 {
		if v, ok := weight[[2]string{start, end}]; ok {
			return v
		}
		if _, ok := graph[start]; !ok {
			return -1.0
		}
		if _, ok := graph[end]; !ok {
			return -1.0
		}
		if _, ok := visited[start]; ok {
			return -1.0
		}

		visited[start] = true
		r := 0.0
		for _, v := range graph[start] {
			value := _dfs(v, end, visited)
			if value != -1.0 {
				r = value * weight[[2]string{start, v}]
				weight[[2]string{start, end}] = r
				return r
			}
		}

		delete(visited, start)
		return -1.0
	}

	res := make([]float64, 0)
	for _, query := range queries {
		res = append(res, _dfs(query[0], query[1], make(map[string]bool, 0)))
	}
	return res
}
