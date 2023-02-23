package leetcode

func findSolution(customFunction func(int, int) int, z int) [][]int {
    type pair struct{x, y int}
    
    record := make(map[pair]bool)
    
    res := [][]int{}
    
    var find func(x, y int)
    
    find = func(x, y int){
        if customFunction(x, y) >= z{
            if customFunction(x, y) == z{
                res = append(res, append([]int{}, x, y))
                record[pair{x, y}] = true
            }
            return
        }
        
        if !record[pair{x+1, y}]{
            find(x+1, y)
            record[pair{x+1, y}] = true
        }
        
        if !record[pair{x, y+1}]{
            find(x, y+1)
            record[pair{x, y+1}] = true
        }
        
    }
    
    find(1, 1)
    
    return res
}