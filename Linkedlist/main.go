package main

import (
	list "Linkedlist/Linkedlist"
)

func main()  {
	myList := list.New()
	list.Print(myList)

	myList.Append(list.NewNode(50))
	myList.Append(list.NewNode(51))
	myList.Append(list.NewNode(56))
	myList.Append(list.NewNode(57))

	list.Print(myList)
	list.PrintInv(myList)
}