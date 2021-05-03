package apt

import "math"

type Node interface {
	Eval(x, y float32) float32
	String() string
}

// Leaf Node (0 Children)
type LeafNode struct{}

// Single Node (sin/cos)
type SingleNode struct {
	Child Node
}

// Double Node (+,-)
type DoubleNode struct {
	LeftChild  Node
	RightChild Node
}

// X
type OpX LeafNode

func (op *OpX) Eval(x, y float32) float32 {
	return x
}

func (op *OpX) String() string {
	return "X"
}

// Y
type OpY LeafNode

func (op *OpY) Eval(x, y float32) float32 {
	return y
}

func (op *OpY) String() string {
	return "Y"
}

// PLus
type OpPlus DoubleNode

func (op *OpPlus) Eval(x, y float32) float32 {
	return op.LeftChild.Eval(x, y) + op.RightChild.Eval(x, y)
}

func (op *OpPlus) String() string {
	return "( + " + op.LeftChild.String() + " " + op.RightChild.String() + ")"
}

// Sinus
type OpSin SingleNode

func (op *OpSin) Eval(x, y float32) float32 {
	return float32(math.Sin(float64(op.Child.Eval(x, y))))
}

func (op *OpSin) String() string {
	return "( Sin " + op.Child.String() + " )"
}
