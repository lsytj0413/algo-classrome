#[allow(dead_code)]
pub struct Solution {}

use std::collections::HashMap;

impl Solution {
    #[allow(dead_code)]
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut record = HashMap::with_capacity(nums.len());
        for (index, num) in nums.iter().enumerate() {
            let found = record.contains_key(&(target-num));
            if !found {
                record.insert(num, index);
            } else {
                return vec![*record.get(&(target-num)).unwrap() as i32, index as i32]
            }
        }
        vec![]
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_two_sum(){
        assert_eq!(vec![0,1], Solution::two_sum(vec![2,7,11,15], 9));
        assert_eq!(vec![1,2], Solution::two_sum(vec![3,2,4], 6));
    }
}