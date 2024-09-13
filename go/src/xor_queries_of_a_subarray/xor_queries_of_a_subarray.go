package leetcode


func xorQueries(arr []int, queries [][]int) []int {
	n := len(arr)
	prefix := make([]int, n)
	prefix[0] = arr[0]
	out := make([]int, len(queries))
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] ^ arr[i]
	}
	for i, q := range queries {
		start := q[0]
		end := q[1]
		res := prefix[end]
		if start > 0 {
			res ^= prefix[start-1]
		}
		out[i] = res
	}
	return out
}
