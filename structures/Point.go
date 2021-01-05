package structures

// Point defines a two-dimensional coordinate point
type Point struct {
	X, Y int
}

// Intss2Points converts [][]int to []Point
func Intss2Points(points [][]int) []Point {
	res := make([]Point, 0, len(points))
	for i, p := range points {
		res[i] = Point{X: p[0], Y: p[1]}
	}

	return res
}

// Points2Intss converts []Point to 　[][]int
func Points2Intss(points []Point) [][]int {
	res := make([][]int, 0, len(points))

	for i, p := range points {
		res[i] = []int{p.X, p.Y}
	}

	return res
}
