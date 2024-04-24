package main

import "fmt"

func main(){

	//calculating the gcd of two integers using EUclidean algorithm
	a := 15
	b := 36
	c := 0
	e := 0

	if a > b {
		c = a 
		e = b
	} else {
		c = b
		e = a
	}

	for i:= 1; i <= c; i++ {
			if i*e + c%i == c {
			fmt.Printf("the remainder is:%d, the divisor is %d", i,e)
		}
	}
}

// 