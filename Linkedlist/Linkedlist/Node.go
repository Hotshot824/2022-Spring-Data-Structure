package Linkedlist

type Node struct {
	data interface{}
	Prev *Node
	Next *Node
}

type Linkedlist struct{
	head *Node
	tail *Node
}

func New() *Linkedlist {
	return new(Linkedlist)
}

func NewNode (Val interface{}) *Node{
	node := new(Node)
	node.data = Val
	return node
}

func (list *Linkedlist) Head() *Node{
    return list.head
}

func (list *Linkedlist) Tail() *Node{
    return list.tail
}

func (list *Linkedlist) isEmpty() bool{
	return list.head == nil
}

func (list *Linkedlist) Prepend(node *Node){
	if list.isEmpty(){
		list.head = node
		list.tail = node
		return
	}
	list.head.Prev = node.Next
	node.Next = list.head
	list.head = node
}

func (list *Linkedlist) Append(node *Node){
    if list.isEmpty(){
        list.head = node
        list.tail = node
        return
    }
	node.Prev = list.tail
    list.tail.Next = node
    list.tail = node  
}