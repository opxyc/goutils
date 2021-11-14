package main

import (
	"context"
	"fmt"

	"github.com/opxyc/goutils/sch"
)

func main() {
	ch := make(chan struct{})
	sch.PingAt(context.Background(), 18, 53, 00, ch)
	for {
		<-ch
		fmt.Println("main// pinged")
	}
}
