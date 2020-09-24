package stack

type (
	node struct {
		data interface{}
		prev *node
	}
	Stack struct {
		tail *node
		len  int
	}
)

func New() *Stack {
	return &Stack{
		tail: nil,
		len:  0,
	}
}

func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) Empty() bool {
	return s.len == 0
}

func (s Stack) Top() interface{} {
	if s.len == 0 {
		return nil
	}
	return s.tail.data
}

func (s *Stack) Push(data interface{}) {
	n := &node{
		data: data,
		prev: s.tail,
	}
	s.tail = n
	s.len++
}

func (s *Stack) Pop() interface{} {
	if s.len == 0 {
		return nil
	}
	n := s.tail
	s.tail = n.prev
	s.len--
	return n.data
}
