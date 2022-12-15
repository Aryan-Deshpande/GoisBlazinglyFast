package concurrency

import(
	"fmt"
	"sync"

)

func waitg(){
	var wg sync.WaitGroup
	wg.Add(10)

	go calling(1)
	go calling(2)


}

func calling(i int){
	for i:=0; i< 10; i++{
		if(i== 1){
			fmt.Println("1")
		}else{
			fmt.Println("2")
		}
	}
}