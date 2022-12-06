package main

import (
	"fmt"
)

func main(){
	// slices
	var a = []int{1,2,3,4,5,6,7};
	fmt.Println(a);

	for i:=0; i<len(a); i++{
		fmt.Println(a[i]);
	}
}