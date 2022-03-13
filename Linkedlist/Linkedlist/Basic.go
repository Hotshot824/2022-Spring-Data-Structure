package Linkedlist

import (
	"fmt"
)

func Print(list *Linkedlist){
    for current := list.Head(); current != nil; current = current.Next{
        fmt.Printf("%v -> ", current.data)
    }
    fmt.Println("nil")
}

func PrintInv(list *Linkedlist){
    for current := list.Tail(); current != nil; current = current.Prev{
        fmt.Printf("%v -> ", current.data)
    }
    fmt.Println("nil")
}