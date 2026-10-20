package main

import (
	"fmt"
	"os"
)

var File string
var Linktype uint32

func main() {
	File = os.Args[1]
	fmt.Printf("Global PCAP Header: --\n\n")
	head()
	fmt.Printf("\n\nSome shit about the first packet idk: --\n\n")
	pack()
}
