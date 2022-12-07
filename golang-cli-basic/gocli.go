package main

import (
	"fmt"
	"os"
	"time"
	//"os"
	//"io"
)

func main(){
	var a int = 3;
	var b int = 4;
	fmt.Println("enter value for a")
	fmt.Scan(&a);
	fmt.Println("enter value for b")
	fmt.Scan(&a);
	var c bool = transfer_details(a,b);
	fmt.Println(c);
}

func transfer_details(x int,y int) bool {
	fmt.Print("transfering details from ", x, " to ", y)
	fmt.Print("... ..... ..... ..... ..... ..... ");
	time.Sleep(3 * 4);
	var a int = 0;
	fmt.Scan(&a);

	if( a!=1 || a != 0){
		fmt.Print("error");
		os.Exit(1);

	}else if(a == 1){
		fmt.Print("success");
		return true;
	}else if(a==0){
		os.Exit(1);
	}else{
		fmt.Print("joe");
	}
	return false;
}