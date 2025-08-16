package main
import "fmt"

func main(){
    // long way
    var name string = "John"
    // short way and multiple declation
    age,isDev := 21,true
    // const
    const G = 9.8
    
    fmt.Println(name, age, isDev, G)
}