package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	_, err := os.Stat("./basics/hello.txt");

	if err == nil {
		readthefile("./basics/hello.txt");
		fmt.Println("hello there");

		} else {

		

		file, err := os.Create("./basics/hello.txt");
		if err !=nil {
			panic(err);
		}

		content := "hey"

		length, err := io.WriteString(file,content);

		defer file.Close()

		fmt.Println(length)
	}
	
}

func readthefile(filename string){
	bytes, err := ioutil.ReadFile(filename);

	if err != nil {
		fmt.Println(err);
	}

	fmt.Println(bytes, "these are the bytes")
}