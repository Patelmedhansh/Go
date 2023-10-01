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
// ALSO 

/*length := getLength(email)
if length < 1 {
    fmt.Println("Email is invalid")
}
can be declaired as 

if length := getLength(email); length < 1 {
    fmt.Println("Email is invalid")
}
*/ 




/*Here are some of the comparison operators in Go:

== equal to
!= not equal to
< less than
> greater than
<= less than or equal to
>= greater than or equal to*/ 