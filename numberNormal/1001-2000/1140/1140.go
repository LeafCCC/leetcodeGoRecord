package leetcode

func stoneGameII(piles []int) int {
    //dfs(i, m)
    //后缀和 s[i]
    //dfs(i, m) = s[i]- min( dfs(i+x, max(x,m)) )( 1<=x<=2*m)
    
    n := len(piles)
    cache := make([][]int, n)
    for i := range cache{
        cache[i] = make([]int, n)
        for j := range cache[i]{
            cache[i][j] = -1
        }
    }
    
    //后缀和
    s := make([]int, n)
    s[n-1] = piles[n-1]
    for i := n - 2; i >= 0; i--{
        s[i] = s[i+1] + piles[i] 
    }
    
    var dfs func(int, int)int
    
    dfs = func(i, m int) int{
        if i + 2*m >= n{
            return s[i]
        }
        
        if cache[i][m] != -1{
            return cache[i][m]
        }
        
        tmp := dfs(i+1, max(1,m))
        
        for x := 2; x <= 2*m; x++{
            tmp = min(dfs(i+x, max(x, m)), tmp)
        }
        
        cache[i][m] = s[i] - tmp
        
        return cache[i][m]
    }
    
    return dfs(0, 1)
}

func min(a, b int) int{if a < b{return a};return b}
func max(a, b int) int{if a < b{return b};return a}