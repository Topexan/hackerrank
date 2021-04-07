func arrayManipulation(n int32, queries [][]int32) int64 {
    var max int64 = 0
    var arr []int32 = make([]int32, n)
    for i := range arr {
        arr[i] = 0
    }
    for i := 0; i < len(queries); i++ {
        for j := queries[i][0]; j <= queries[i][1]; j++ {
            arr[j-1] = arr[j-1] + queries[i][2]
            //fmt.Println(j)
        }
    }
    for i := range arr {
        if int64(arr[i]) > max {
            max = int64(arr[i])
        }
    }
    return max
}