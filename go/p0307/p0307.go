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

// type NumArray struct {
// 	arr []int
// }

// func Constructor(nums []int) NumArray {
// 	arr := make([]int, len(nums))
// 	copy(arr, nums)
// 	return NumArray{
// 		arr: arr,
// 	}
// }

// func (this *NumArray) Update(i int, val int) {
// 	this.arr[i] = val
// }

// func (this *NumArray) SumRange(i int, j int) int {
// 	sum := 0
// 	for ; i <= j; i++ {
// 		sum += this.arr[i]
// 	}
// 	return sum
// }

// DOC: 树状数组解法
type NumArray struct {
	arr  []int
	sums []int
}

func lowbit(x int) int {
	return x & (0 - x)
}

func change(i int, val int, arr []int) {
	for i < len(arr) {
		arr[i] += val
		i += lowbit(i)
	}
}

func sum(n int, arr []int) int {
	sum := 0
	for n > 0 {
		sum += arr[n]
		n -= lowbit(n)
	}
	return sum
}

func Constructor(nums []int) NumArray {
	arr := make([]int, len(nums))
	copy(arr, nums)

	sums := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		change(i+1, nums[i], sums)
	}
	return NumArray{
		arr:  arr,
		sums: sums,
	}
}

func (this *NumArray) Update(i int, val int) {
	delta := val - this.arr[i]
	this.arr[i] = val
	change(i+1, delta, this.sums)
}

func (this *NumArray) SumRange(i int, j int) int {
	return sum(j+1, this.sums) - sum(i, this.sums)
}

// TODO: 线段树解法

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
