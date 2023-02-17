package leetcode

func multiply(num1 string, num2 string) string {
    m, n := len(num1), len(num2)
    ans := make([]int, m+n) //结果最大为m+n位
    
    for i := 0; i < m; i++{
        for j := 0; j < n; j++{
            t1 := int(num1[m-i-1] - '0')
            t2 := int(num2[n-j-1] - '0')
            tmp := t1*t2
            if tmp < 10{
                ans[i+j] += tmp
            }else{
                ans[i+j] += tmp%10
                ans[i+j+1] += tmp/10
            }
            
            if ans[i+j] >= 10{
                ans[i+j] %= 10
                ans[i+j+1]++
            }
            
            if ans[i+j+1] >= 10{
                ans[i+j+1] %= 10
                ans[i+j+2]++
            }
        }
    }
    
    for ans[len(ans)-1] == 0 && len(ans) > 1{
        ans = ans[:len(ans)-1]
    }
    
    res := []byte{}
    for i := len(ans) - 1; i >= 0; i--{
        res = append(res, byte(ans[i] + '0'))
    }
    return string(res)
    
}