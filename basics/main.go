package main

import (
	"fmt"
	//"strings"
	//"strconv"
	"basics/passby" // here we are importing a local package, by mentioning the package name in main.go
)


func main(){
	fmt.Println("Main go")
	passby.V(); // we are calling the package function via packagename.functioncall()
}
