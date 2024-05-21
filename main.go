package gopackage

import (
	"fmt"
	"log"

	
)

func LogInfo(msg string){
	log.Printf("INFO- %v",msg)
}

func LogWarn(msg string){
	log.Printf("WARN- %v",msg)
} 

// //erc20 like functions
// func Add(a int,b int)int{
// sum:= a+b
// return sum
// }


// add two number checking for overflow
func Add(b int, q int) (int, error) {
 
    // Check overflow
    var sum int = q + b
 
    if (sum < q || sum < b) == (b >= 0 && q >= 0) {
        return 0, fmt.Errorf("math: addition overflow occurred %d + %d", b, q)
    }
 
    return sum, nil
}