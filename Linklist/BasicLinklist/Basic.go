package BasicLinklist

type Node struct {
	Val interface{}
	prev *Node
}

type Linkedlist struct{
	head *Node
	tail *Node
}

func New() *Linkedlist {
	return new(Linkedlist)
}

func NewNode (Val interface{}) *Node{
	n := new(Node)
	n.Val = Val
	return n
}