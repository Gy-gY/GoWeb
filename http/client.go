﻿package main
package http
package 1228
package 1228-1
package 1228-2
package 1229-core-1002

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server")
		os.Exit(1)
	}
	serverAddress := os.Args[1]

	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	
	args := Args{17, 123}
	var reply int
	err = client.Call("Math.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Math error:", err)
	}
	fmt.Printf("Math: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Math.Divide", args, &quot)
	if err != nil {
		log.Fatal("Math error:", err)
	}
	fmt.Printf("Math: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

}
