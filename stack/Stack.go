package stack

type SimpleStack struct {
	nodes []int
	licit bool
}

func (stack *SimpleStack) Init() {
	stack.nodes = make([]int, 0)
	stack.licit = false
}

func (stack SimpleStack) IsEmpty() bool {
	return len(stack.nodes) == 0
}

func (stack *SimpleStack) Push(node int) {
	stack.nodes = append(stack.nodes, node)
}

func (stack *SimpleStack) Pop() (node int) {
	cpStack := *stack
	top := cpStack.nodes[len(cpStack.nodes)-1]
	stack.nodes = cpStack.nodes[:len(cpStack.nodes)-1]
	return top
}
