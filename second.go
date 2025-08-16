// Day 2: Data Types & Input/Output
package main
import "fmt"

func main(){
    // string
    var name string
    
    // integer
    var age int
    
    // float
    var height float64
    
    // boolean
    var isDev bool
    
    // Taking inputs from user
    fmt.Println("Enter your name: ")
    fmt.Scanln(&name)
    
    fmt.Println("Enter your age (INT): ")
    fmt.Scanln(&age)
    
    fmt.Println("Enter your height (DECIMAL): ")
    fmt.Scanln(&height)
    
    fmt.Println("Are your a developer (true/false)? ")
    fmt.Scanln(&isDev)

    // NOTE: WE CAN ALSO USE fmt.Scanf() FOR TAKING INPUT
    
    // showing outputs:
    fmt.Printf("Name: %s, Age: %d, Height: %.2f, Developer: %t",name,age,height,isDev) 
}