package concurrency

import (
	"fmt"
	//"sync"
	"time"
)

func chann(){ // channel is way to communicate w/
				// each other

	// we can use *make* to create channels, slices, maps
	c := make(chan string);
	go count("sheep", c); 
	msg, status := <- c; // status is a secondary output // shows the status channel ( close or open )

	// channels can be used to communicate between go routines // or sender/recvr

	if(!status){
		fmt.Println(status);
	}
	fmt.Println(msg)

}

func count(things string, c chan string){
	for i := 0; i < 5; i++ {
		c <- things
		time.Sleep(time.Millisecond * 500)
	}
	close(c); // use close do avoid deadlock
	// deadlock occurs in the recvr when, sender work is done
	// but the recvr waits until another msg is sent
}


