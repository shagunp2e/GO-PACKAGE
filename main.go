package gopackage

import (
	// "fmt"
	"log"

	
)

func LogInfo(msg string){
	log.Printf("INFO- %v",msg)
}

func LogWarn(msg string){
	log.Printf("WARN- %v",msg)
} 

//erc20 like functions
func Add(a int,b int)int{
sum:= a+b
return sum
}

