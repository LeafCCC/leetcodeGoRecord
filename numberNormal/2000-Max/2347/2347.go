package leetcode

func bestHand(ranks []int, suits []byte) string {
    sort.Slice(suits, func(i, j int) bool{
        return suits[i] < suits[j]
    })
    
    if suits[0] == suits[4]{
        return "Flush"
    }
    
    record := make(map[int]int)
    for _, val := range ranks{record[val]++}
    
    max := 0
    for _, val := range record{
        if max < val{
            max = val
        }
    }
    
    switch max{
        case 3,4,5:
            return "Three of a Kind"
        case 2:
            return "Pair"
    }
    
    return "High Card"
}