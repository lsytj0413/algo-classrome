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

package p0468

import (
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	validIPv4 := func(IP string) bool {
		ips := strings.Split(IP, ".")
		if len(ips) != 4 {
			return false
		}

		for _, v := range ips {
			if len(v) == 0 || (len(v) > 1 && (v[0] == '0' || v[0] == '-')) {
				return false
			}

			n, err := strconv.Atoi(v)
			if err != nil || n < 0 || n >= 256 {
				return false
			}
		}
		return true
	}
	validIPv6 := func(IP string) bool {
		ips := strings.Split(IP, ":")
		if len(ips) != 8 {
			return false
		}

		for _, v := range ips {
			if len(v) == 0 || (len(v) > 1 && (v[0] == '-')) || len(v) > 4 {
				return false
			}

			n, err := strconv.ParseInt(v, 16, 0)
			if err != nil || n < 0 || n >= 65536 {
				return false
			}
			if n == 0 && len(v) > 4 {
				return false
			}
		}
		return true
	}

	switch {
	case strings.Index(IP, ".") != -1 && validIPv4(IP):
		return "IPv4"
	case strings.Index(IP, ":") != -1 && validIPv6(IP):
		return "IPv6"
	}
	return "Neither"
}
