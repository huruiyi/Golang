package main

import (
	"fmt"

	"../../goflow"
	flow "github.com/trustmaster/goflow" /*vo分支*/
)

func main() {

	net := goflow.NewEncodingApp()

	in := make(chan string)
	net.SetInPort("In", in)

	flow.RunNet(net)

	for i := 0; i < 20; i++ {
		in <- fmt.Sprint("Message", i)
	}

	close(in)
	<-net.Wait()
}
