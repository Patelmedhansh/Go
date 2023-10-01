package main 
import "fmt"
func main(){

	var height int
	

	fmt.Println("Enter your Height in Ft") 
	fmt.Scanln(&height) 


	if height > 5{
		fmt.Println("You are very tall !")
	} else if height >=3 {
		fmt.Println("You are average height ") 
	} else {
		fmt.Println("Hey Shawty !!!")
 	}
	

}



/*Here are some of the comparison operators in Go:

== equal to
!= not equal to
< less than
> greater than
<= less than or equal to
>= greater than or equal to*/ 