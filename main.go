package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/David-Antunes/network-emulation-proxy/xdp"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcapgo"
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
	conn, err := net.Dial("unix", os.Args[1])

	if err != nil {
		panic(err)
	}
	dec := gob.NewDecoder(conn)
        f, err := os.Create("temp.pcap")
        if err != nil {
                log.Fatal(err)
                os.Exit(1)
        }
        defer f.Close()

        w := pcapgo.NewWriter(os.Stdout)
        w.WriteFileHeader(uint32(65535), layers.LinkTypeEthernet)
	for {
		var frame *xdp.Frame

		err := dec.Decode(&frame)
		if err != nil {
			panic(err)
		}
		packet := gopacket.NewPacket(frame.FramePointer, layers.LinkTypeEthernet, gopacket.Default)
        w.WritePacket(gopacket.CaptureInfo{Timestamp: frame.Time, CaptureLength: frame.FrameSize, Length: frame.FrameSize,InterfaceIndex: 1}, packet.Data())
	}

}
