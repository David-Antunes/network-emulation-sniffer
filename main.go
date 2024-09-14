package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"

	"github.com/David-Antunes/network-emulation-proxy/xdp"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("missing socket id")
		return 
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments")
		return 
	}
	conn, err := net.Dial("unix", "/tmp/"+os.Args[1]+".sniff")

	if err != nil {
		panic(err)
	}
	dec := gob.NewDecoder(conn)
	for {
		var frame *xdp.Frame

		err := dec.Decode(&frame)
		if err != nil {
			panic(err)
		}
		fmt.Println(frame.FramePointer)
	}

}
