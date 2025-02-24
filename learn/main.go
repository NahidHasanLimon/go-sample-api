package main

import (
    "fmt"
	"learn/learnpkg"
	"strings"
)

func main() {
    fmt.Println("Hy from SayHello Function")
	fmt.Println(learn.SayHello("Hy from SayHello Function"))
   
	r:= strings.NewReader(" A New Reader Initiated")
	buff:= make([]byte, 4)
	n,err :=r.Read(buff)
	fmt.Println(buff[:4],err)
	fmt.Println(n)


	// for{
		

	// 	fmt.Println(buff[:n], err)
	// 	if(err != nil){
	// 		fmt.Println("breaking")
	// 		break;
	// 	}

		
	// }


}