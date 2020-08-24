package main

import (
	"fmt"
	"os"
)

var (
	articlesNo 			int16
	impactFactor 		int16
	scientistsToBribe	int16

)

func readInput() {
	n, err	:= fmt.Scanf("%d %d\n",&articlesNo, &impactFactor)

	if( err != nil ) {
		fmt.Fprintln(os.Stderr,"error:" +err.Error())

	} else if( n != 2 ) {
		fmt.Fprintf(os.Stderr, "expecting exactly 2 inputs, got %d\n",n)
	}
}

func compute() int16 {
	/**
	*	impactFactor = (scientistsToBribe) / (articlesNo)
	*/

	// Case 1 : If articlesNo (divisor) = 1
	if (articlesNo == 1) {
		return impactFactor
	}
	// Case 2 : Otherwise
	return (articlesNo * (impactFactor-1))+1

}

func main() {
	readInput()
	fmt.Printf("%d\n",compute())
}