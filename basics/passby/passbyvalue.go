package passby

import (
	"fmt"
)

func V(){
	var no2 int = 5;
	ptr := &no2;

	temp:= change2(ptr);
	fmt.Println(temp);
}

func change2(a *int) *int{
	*a = 12;
	return a;
}