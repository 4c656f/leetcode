package leetcode

import "sort"

type Pt struct{
    val int
    isEnd int
}
func minGroups(intervals [][]int) int {
    pts := make([]Pt,0 , len(intervals)*2)
    for _, n := range intervals{
        pts = append(pts, Pt{n[0], 1})
        pts = append(pts, Pt{n[1], 0})
    }

    sort.Slice(pts, func(i,j int)bool{
        if pts[i].val == pts[j].val{
            return pts[i].isEnd > pts[j].isEnd
        }
        return pts[i].val < pts[j].val
    })
    max := 0
    cur := 0

    for _, pt := range pts{
        if pt.isEnd == 1{
            cur++
        }else{
            cur--
        }
        if cur > max{
            max = cur
        }
    }


    return max
}