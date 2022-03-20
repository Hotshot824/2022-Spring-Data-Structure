package main

import (
	"Stack/Stack"
	"fmt"
)

func main(){
	var choice int = 0
	for{
		fmt.Print("Choice : \t1.Inputdata\t2.Showdata\n")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			Stack.InputData()
		case 2:
			Stack.ShowData()
		}
	}
}