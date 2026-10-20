package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func pack() {
	file, err := os.Open(File)
	if err != nil {
		panic(err) // tells us what shit happened
	}

	defer file.Close() // i still dont know wtf defer does

	pack := make([]byte, 40)
	n, err := file.Read(pack)
	if err != nil {
		panic(err) // this will usually be a read error (?)
	}
	fmt.Printf("Read %v bytes from %v.\n", n, File)
	fmt.Printf("Packet Data [24:40]: %x\n", pack[24:40])

	timestamp1 := parse(pack[24:28]) // seconds since the epoch was captured
	timestamp2 := parse(pack[28:32]) // fractional part of the timestamp (microseconds)
	secondsfrac := float64(timestamp2) / 1000000.0
	length := parse(pack[32:36]) // packet length
	oglength := parse(pack[36:40])
	fmt.Printf("BINARY: \n")
	fmt.Printf("Timestamp [24:28]: %x\n", timestamp1)
	fmt.Printf("Timestamp Fraction [28:32]: %x\n", timestamp2)
	fmt.Printf("Packet Length [32:36]: %x\n", length)
	fmt.Printf("Packet Origional Length [36:40]: %x\n", oglength)

	fmt.Printf("NUMERIC:\n")
	fmt.Printf("Timestamp (seconds) [24:28]: %v\n", timestamp1)
	fmt.Printf("Timestamp Fraction (microseconds) [28:32]: %v (%v seconds)\n", timestamp2, secondsfrac)
	fmt.Printf("Packet Length (bytes) [32:36]: %v\n", length)
	fmt.Printf("Packet Origional Length (bytes) [36:40]: %v\n", oglength)

	data := make([]byte, length) // length = 98
	_, err = io.ReadFull(file, data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Packet data (%d bytes):\n%x\n", len(data), data)
	targ := binary.BigEndian.Uint16(data[0:6])
	src := binary.BigEndian.Uint16(data[6:12])
	fmt.Printf("Target MAC addr: %x\nSource MAC addr: %x\n", targ, src)

}
