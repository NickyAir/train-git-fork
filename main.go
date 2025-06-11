package main

import (
	"fmt"
	"errors"
)

func add(a, b int)int{
	return a+b
}
func sub(a, b int)int{
	return a-b
}
func mul(a, b int)int{
	return a*b
}
func div(a, b int)(int, error){
	if b == 0{
		return 0, errors.New("Cannot divide by zero")
	}
	return a/b, nil
}

func main(){
	a := 2
	b := 2

	fmt.Printf("Add: %d\n", add(a, b))
	fmt.Printf("Sub: %d\n", sub(a, b))
	fmt.Printf("Mul: %d\n", mul(a, b))
	
	res, err := div(a, b)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("Div: %d\n", res)
	}
}