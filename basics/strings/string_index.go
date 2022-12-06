package strings

import(
	"fmt"
	"strings"
	"strconv"
)

func stringm(){
	var a string = "hello there"
	// 1.0 extract particular index
	fmt.Println(strings.ContainsAny(a, "xt")); // returns true if any one substring exists in the original string

	// 2.0
	fmt.Println(strings.Index(a,"h")); // prints the index at which the substring is present from the original string

	// 3.0
	x:= "hey " 
	name:= ""
	
	fmt.Scanln(&name);
	fmt.Println(x+name) // reads cli input, to print custom output
	
	// 4.0 // replacing first few instance of string based on // original string, substring to be replaced, new substring to replace and no. of repl
	z:= "joe joe joe is joe"
	zz:="doe"
	zzz:= strings.Replace(z, "joe", zz, 2)
	fmt.Println(zzz)

	// 5.0 // int to string & string to int
	string1:="1"

	num1, err:= strconv.Atoi(string1)
	if err!=nil{
		panic(err)
	}else{
		fmt.Println(num1);
	}
}