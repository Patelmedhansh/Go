package main 
import "fmt"

func concat (s1 string , s2 string ) string{
	return s1 + s2 

}

func main (){

	/* func (x int , y int )int {
		return x-y
	} */

	// the above is function signature, it basically states what funciton does and describes the input and output 

fmt.Println(concat("HEY", "CHEEMKI"))
fmt.Println(concat("cheemki is ", "chibavli ")) 


}