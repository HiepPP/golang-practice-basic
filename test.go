package main

import "fmt"

func main(){
	var p *int
	fmt.Printf("%v", p)

	var i int
	p = &i
	*p = 8
	fmt.Printf("%v", i)
	//fmt.Printf("%v", p)
}

