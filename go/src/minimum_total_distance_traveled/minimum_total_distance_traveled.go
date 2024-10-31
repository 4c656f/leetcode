package leetcode

import (
	"math"
	"sort"
)
func minimumTotalDistance(robot []int, factory [][]int) int64 {
    dp := make([][]int64, len(robot))
    n := len(robot)
    k := len(factory)
    
    sort.Ints(robot)
    factories := make([]int, 0, k)
    sort.Slice(factory, func(i,j int)bool{
        return factory[i][0] < factory[j][0]
    })
    for _, f := range factory{
        for i := 0; i < f[1]; i++{
            factories = append(factories, f[0])
        }
    }
    k = len(factories)
    for i := range n{
        dp[i] = make([]int64, k)
        for j := range k{
            dp[i][j] = -1
        }
    }
    var dfs func(i, j int)int64

    dfs = func(i, j int)int64{
        if i == -1{
            return 0
        }
        if j == -1 {
            return math.MaxInt64
        }
        if dp[i][j] != -1{
            return dp[i][j]
        }
        var res int64 = dfs(i, j-1)
        ass := dfs(i-1, j -1)
        if ass < res{
            sm := ass + abs(int64(robot[i] - factories[j]))
            res = min(res, sm)
        }
        dp[i][j] = res
        return res
    }

    return dfs(n -1, k -1)
}

func min(n1, n2 int64)int64{
    if n1 < n2{
        return n1
    }
    return n2
}

func abs(n int64)int64{
    if n < 0{
        return -n
    }
    return n
}