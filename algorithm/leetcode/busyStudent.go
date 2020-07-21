package leetcode

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	result := 0
	for k,s := range startTime{
		if s > queryTime{
			continue
		}
		if endTime[k] >= queryTime{
			result++
		}
	}
	return result
}

func coinChange(coins []int, amount int) int {

}