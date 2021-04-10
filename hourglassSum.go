package hackerrank

func hourglassSum(arr [][]int32) int32 {
	var res int32 = -99999
	var sum int32
	var sums []int32
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum = arr[j][i] + arr[j][i+1] + arr[j][i+2] + arr[j+1][i+1] + arr[j+2][i] + arr[j+2][i+1] + arr[j+2][i+2]
			sums = append(sums, sum)
		}
	}
	for i := range sums {
		if sums[i] > res {
			res = sums[i]
		}
	}
	return res
}
