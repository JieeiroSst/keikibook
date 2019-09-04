package algorithm

type Node struct {
	Value int
}

//stack algorithm
type Stack struct {
	node  []*Node
	count int64
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(n *Node) {
	s.node = append(s.node[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.node[s.count]
}

//Queue algorithm
type Queue struct {
	node  []*Node
	size  int
	head  int
	tail  int
	count int
}

func (q *Queue) Push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		node := make([]*Node, len(q.node)+q.size)
		copy(node, q.node[q.head:])
		copy(node[len(q.node)-q.head:], q.node[:q.head])
		q.head = 0
		q.tail = len(q.node)
	}
	q.node[q.tail] = n
	q.tail = (q.tail + 1) % len(q.node)
	q.count++
}

func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.node[q.head]
	q.head = (q.head + 1) % len(q.node)
	q.count--
	return node
}
