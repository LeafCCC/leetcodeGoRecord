package leetcode

func rotate(matrix [][]int)  {
    n := len(matrix)
    for i := 0; i < n/2; i++{
        begin, end := i, n-i-1
        for j := begin; j < end; j++{
            matrix[i][j], matrix[i+j-begin][end], matrix[n-i-1][end-j+begin], matrix[n-i-1-j+begin][begin] = matrix[n-i-1-j+begin][begin],  matrix[i][j], matrix[i+j-begin][end], matrix[n-i-1][end-j+begin]
        }
    }
}