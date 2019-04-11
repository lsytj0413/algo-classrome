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

package p0721

import (
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	type SetsID struct {
		name string
		id   int
	}

	var id int
	emailTrees := make(map[string]SetsID)
	accountTrees := make(map[SetsID][]string)

	for _, account := range accounts {
		name := account[0]
		emails := account[1:]
		ids := make([]SetsID, 0)
		for _, email := range emails {
			if r, ok := emailTrees[email]; ok {
				ids = append(ids, r)
			}
		}

		if len(ids) == 0 {
			nid := SetsID{
				name: name,
				id:   id,
			}
			id++
			for _, email := range emails {
				emailTrees[email] = nid
			}
			accountTrees[nid] = emails
			continue
		}

		firstid := ids[0]
		for _, idv := range ids {
			if idv.id < firstid.id {
				firstid = idv
			}
		}

		for _, idv := range ids {
			idvemails := accountTrees[idv]
			delete(accountTrees, idv)
			emails = append(emails, idvemails...)
		}
		accountTrees[firstid] = emails
		for _, email := range emails {
			emailTrees[email] = firstid
		}
	}

	names := make([]SetsID, 0, len(accountTrees))
	for aid := range accountTrees {
		names = append(names, aid)
	}
	sort.Slice(names, func(i int, j int) bool {
		if names[i].name != names[j].name {
			return names[i].name < names[j].name
		}
		return names[i].id < names[j].id
	})

	res := make([][]string, 0, len(names))
	for _, name := range names {
		v := []string{name.name}
		emails := accountTrees[name]
		sort.Strings(emails)
		for i := range emails {
			switch {
			case i == 0:
				v = append(v, emails[i])
			case i > 0:
				if emails[i] != emails[i-1] {
					v = append(v, emails[i])
				}
			}
		}
		res = append(res, v)
	}
	return res
}
