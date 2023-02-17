package leetcode

func permuteUnique(nums []int) [][]int {
    res := [][]int{}
    
    sort.Ints(nums)
    
    var backTrack func(level int)
    
    backTrack = func(level int){
        if level == len(nums) - 1{
            res = append(res, append([]int{}, nums...))
            return 
        }
        
        record := make(map[int]bool)
        
        for i := level; i < len(nums); i++{
            if record[nums[i]]{
                continue
            }
            record[nums[i]] = true
            nums[i], nums[level] = nums[level], nums[i]
            backTrack(level+1)
            nums[i], nums[level] = nums[level], nums[i]
        }
    }
    
    backTrack(0)
    
    return res
}