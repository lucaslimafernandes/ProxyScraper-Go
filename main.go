package main

import (
	"fmt"
	"proxyscraper-go/lists"
	"proxyscraper-go/verifier"
)

func main() {

	fmt.Println("Proxyscrap-go")

	p := lists.HandlerProxyscrape()

	verifier.Verifier(p[1])
	// for i := 0; i < 10; i++ {

	// 	verifier.Verifier(p[i])
	// }
}
