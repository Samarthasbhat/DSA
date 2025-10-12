func topKFrequent(nums []int, k int) []int {
    freq := make(map[int]int)
    for _, num := range nums{
        freq[num]++
    }

    n := len(nums)
    buckets := make([][]int, n+1) 
   
   for num, count := range freq{
    buckets[count] = append(buckets[count], num)
   }

   result := []int{}
   for i:= n; i>=0 && len(result) < k; i-- {
    if len(buckets[i]) >0 {
        result = append(result, buckets[i]...)
    }
   }

   if len(result) > k{
    result = result[:k]
   }
   return result
}