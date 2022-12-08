package arrayz

import(
	"fmt"
)

type somethin struct{
	item string
	rating int
} 

func update(){
	var object somethin
	object.item = "iphone"
	object.rating = 10

	ptr:= &object
	exter(ptr)

	fmt.Println(object.rating)
}

func exter(ptr *somethin){
	ptr.rating = 11
}
func up(ptr *somethin) (rating int){
	ptr.rating = 12
	return ptr.rating
}