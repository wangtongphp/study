  package main
  
  import (
      "fmt"
  )
  
  func main() {
      res := twoSum([]int{3, 2, 4}, 6)
      fmt.Println(res)
  }
  
  func twoSum(nums []int, target int) []int {
  
      //键值翻转
      nums_keys := make(map[int][]int, len(nums))
      for k, v := range nums {
          nums_keys[v] = append(nums_keys[v], k)
      }   
  
      // 寻找另一个数是否存在
      for k, v := range nums_keys {
          key := target - k 
  
          if k == key {
              if len(v) > 1 { 
                  return []int{v[0], v[1]}
              }   
          } else if index, ok := nums_keys[key]; ok {
              return []int{v[0], index[0]}
          }   
      }   
  
      return []int{}
  }
  
  /*
  // 可以用二分优化取i2
  func twoSum(nums []int, target int) []int {
      start := 0
      end := sort.IntSlice(nums).Len()
  
      //键值翻转
      nums_keys := make(map[int]int, end)
      for k, v := range nums {
          nums_keys[v] = k
      }
  
      sort.IntSlice(nums).Sort()
      for i := start; i < end-1; i++ {
          for i2 := start + 1; i2 < end; i2++ {
              if nums[i]+nums[i2] > target {
                  start++
                  continue
              } else if nums[i]+nums[i2] == target {
                  return []int{nums_keys[nums[i]], nums_keys[nums[i2]]}
              }
          }
      }
      return []int{}
  }
  */
