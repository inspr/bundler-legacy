package operator

type Operator interface {
	Create()
	Run()
}

type Node interface {
	// Return the list of children nodes
	Next() []Node

	// Run the tasks necessary for the node
	Apply(op ...Operator)
}

type node struct {
	Operators []Operator
}
