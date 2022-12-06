package main

import "fmt"

func main(){
	// printing variables of type int,string,bool,float32
	var a int = 32;
	var b string = "hey there buddio";

	var c bool = true;
	var d float32 = 3.446465;

	// trying arrays
	
	// here there aare two ways of representing an array in golang 
	// -> one with = and var, the other with a := and w/o a var.

	var aa = [100] int{2,3,4,5,5,6,7,7,8,6}; // special case, here [...] is ellipses, although if the rest of indexes are not filled zeros appear.

	bb := [2] string{"hello there", " this is bob"};

	fmt.Println(a,b,c,d,aa,bb[0]);

}