package leetcode

func circularPermutation(n int, start int) []int {
    res := []int{start}
    
    for i := 0; i < n; i++{
        for j := len(res) - 1; j >= 0; j--{
            res = append(res, ((res[j]^start)|(1<<i))^start )
        }
    }
    
    return res
}