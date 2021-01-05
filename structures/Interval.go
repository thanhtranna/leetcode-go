package structures

// Interval provides interval representation
type Interval struct {
	Start int
	End   int
}

// Interval2Ints convert Interval to integer slice
func Interval2Ints(i Interval) []int {
	return []int{i.Start, i.End}
}

// IntervalSlice2Intss converts []Interval to [][]int
func IntervalSlice2Intss(is []Interval) [][]int {
	res := make([][]int, 0, len(is))

	for i := range is {
		res = append(res, Interval2Ints(is[i]))
	}

	return res
}

// Intss2IntervalSlice converts [][]int to []Interval
func Intss2IntervalSlice(intss [][]int) []Interval {
	res := make([]Interval, 0, len(intss))
	for _, ints := range intss {
		res = append(res, Interval{Start: ints[0], End: ints[1]})
	}

	return res
}
