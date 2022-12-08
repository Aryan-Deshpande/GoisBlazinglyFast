package concurrency

import (
	"fmt"
	"time"
	"sync"
)

func Routine(){
	// 1.0 go count("sheep") // these two calls 
	//count("fish") // will run side by side

	// 2.0 defining a wait group
	var w sync.WaitGroup
	w.Add(1); // wait group is a counter that is incremented
	// by w.Add(x)
	// decremented by w.Done()

	go count1("sheep");
	go count1("fish");

	// 3.0 wait group, within anon function
	go func(){
		count1("sheep")
		w.Done() // decrements counter
	}()
	go func(){
		count1("hello")
		w.Done() // decrements counter
	}()

	w.Wait() // blocks/waits till all go routines finished 
}

func count1(things string){
	for i := 0; true; i++ {
		fmt.Println(i, things);
		time.Sleep(time.Millisecond * 600);
	}
}