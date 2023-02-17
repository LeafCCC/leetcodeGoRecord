package leetcode

func groupAnagrams(strs []string) [][]string {
    record := map[string]int{}
    res := [][]string{}
    
    for _, s := range strs{
        tmp := []byte(s)
        sort.Slice(tmp, func(i, j int) bool{return tmp[i] < tmp[j] })
        t := string(tmp)
        
        if record[t] != 0{
            res[record[t] - 1] = append(res[record[t] - 1], s)
        }else{
            res = append(res, []string{s})
            record[t] = len(res)
        }
    }
    
    return res
}