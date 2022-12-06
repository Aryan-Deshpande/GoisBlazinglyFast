package main

import (
	"fmt"
	"reflect"
)

func main(){
	// this is a normal array
	arr := []string {"hey","there","how","are","you"};

	fmt.Println(reflect.TypeOf(arr));

	ans:="" // prints the array

	// to find the length of an array use // len(x) //
	for i := 0; i < len(arr); i++ {

		switch arr[i]{
		case "you":
			ans+=arr[i]+"?"
		case "there":
			ans+=arr[i]+", "
		default:
			ans+=arr[i]+" ";
		}
	}

	fmt.Println(ans) // this prints // "hey there, how are you ?"

}