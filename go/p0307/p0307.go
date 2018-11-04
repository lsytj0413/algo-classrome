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

package p0307

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	arr := make([]int, len(nums))
	copy(arr, nums)
	return NumArray{
		arr: arr,
	}
}

func (this *NumArray) Update(i int, val int) {
	this.arr[i] = val
}

func (this *NumArray) SumRange(i int, j int) int {
	sum := 0
	for ; i <= j; i++ {
		sum += this.arr[i]
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
