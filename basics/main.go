package main

import (
	"fmt"
	//"strings"
	//"strconv"
	//"basics/passby" // here we are importing a local package, by mentioning the package name in main.go
//	"basics/concurrency"
)


// for loop types

func main(){
	v:= []int {1,2,3,4,5,6}
	for _, v := range v { // used to loop over array v
		fmt.Println(v);
	}

	count:= 10
	for i := 0; i < count; i++ { // classic for loop
		fmt.Println(v);
	}

	for{ // infinte for loop that uses 100% CPU
		fmt.Println("hi");
	}

	fmt.Println("Main go") // unreachable becuase of infinite loop

	//concurrency.Routine(); // we are calling the package function via packagename.functioncall()
}
