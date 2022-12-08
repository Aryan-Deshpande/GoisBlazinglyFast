package passby

import (
	"fmt"
)

func R(){
	var number int = 8;

	temp:= change(number);
	
	fmt.Println(temp);
}

func change(no int) int{
	no = 3;
	return no;
}