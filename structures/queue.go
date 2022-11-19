package structures

// Queue Is the queue used to store int
type Queue struct {
	nums []int
}

// NewQueue return *kit.Queue
func NewQueue() *Queue {
	return &Queue{nums: []int{}}
}

// Push n into the queue
func (q *Queue) Push(n int) {
	q.nums = append(q.nums, n)
}

// Pop Take the value that entered the queue first from queue
func (q *Queue) Pop() int {
	res := q.nums[0]
	q.nums = q.nums[1:]
	return res
}

// Len Return the length of queue
func (q *Queue) Len() int {
	return len(q.nums)
}

// IsEmpty Return whether queue is empty
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
