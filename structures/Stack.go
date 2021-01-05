package structures

// Stack is the stack used to store int
type Stack struct {
	nums []int
}

// NewStack returns *kit.Stack
func NewStack() *Stack {
	return &Stack{nums: []int{}}
}

// Push puts n on the stack
func (s *Stack) Push(n int) {
	s.nums = append(s.nums, n)
}

// Pop takes the last value put on the stack from s
func (s *Stack) Pop() int {
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}

// Len returns the length of s
func (s *Stack) Len() int {
	return len(s.nums)
}

// IsEmpty feedback whether s is empty
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}
