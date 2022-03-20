package Stack

import "fmt"

const(
	MAXSTACK = 5
)

var stack[MAXSTACK] string
var top int = -1

func push(data string){
	if(top >= MAXSTACK - 1){
		fmt.Print("Stack is full!\n")
	}else{
		top++
		stack[top] = data
	}
}

func pop() string{
	var data string
	if(top == -1){
		fmt.Print("Stack is empty!\n")
		data = "Null"
		return data
	}else{
		data = stack[top]
		top--
		return data
	}
}

func InputData(){
	fmt.Print("Please enter data : || Enter exit() quit!\n")
	var input string
	for {
		fmt.Scanf("%s", &input)
		if (input != "exit()"){
			push(input)
		} else{
			break
		}
	}
}

func ShowData(){
	for data := pop(); data != "Null"; data = pop(){
		fmt.Print("Pop : ",data, "\n")
	}
}