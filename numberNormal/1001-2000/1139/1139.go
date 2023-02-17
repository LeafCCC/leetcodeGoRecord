package leetcode

func largest1BorderedSquare(grid [][]int) int {
    decide := func(i, j, length int) bool{
        for t := j; t < j + length; t++{
            if grid[i+length-1][t] == 0{
                return false
            }
        }
        
        for t := i+1; t < i + length -1; t++{
            if grid[t][j] == 0 || grid[t][j+length-1] == 0{
                return false
            }
        }
        
        return true
    }
    
    min := func(a,b int) int{
        if a < b{
            return a
        }
        return b
}
    
    res := 0
    
    for i := range grid{
        if len(grid) - i <= res{
            break
        }
        
        now := 0
        tmp := 0
        for now < len(grid[0]){
            if grid[i][now] == 1{
                tmp = now
                for tmp < len(grid[0]) - 1 && grid[i][tmp+1] == 1{
                    tmp++
                }
                for t := min(tmp - now + 1, len(grid) - i); t > res; t--{
                    if decide(i, now, t){
                        res = t
                        break
                    }
                }
            }
            now++
        }
    }
    
    return res*res
}
