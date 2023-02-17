package leetcode

func myPow(x float64, n int) float64 {
    tx, tn := x, n
    if tx < 0{tx = -tx}
    if tn < 0{tn = -tn}
    
    record := []float64{tx}
    i := 2
    for i <= tn{
        record = append(record, record[len(record)-1]*record[len(record)-1])
        i += i
    }

    res := 1.0
    for i := 0; i < len(record); i++{
        if (1<<i)&tn != 0{
            res *= record[i]
        }
    }
    if n < 0{
        res = 1/res
        n = -n
    }
    
    if x < 0 && n % 2 == 1{
        res = -res
    }
    return res
}