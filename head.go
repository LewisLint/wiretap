package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func head() {
	file, err := os.Open(File)
	if err != nil {
		panic(err) // tells us what shit happened
	}

	defer file.Close() // i still dont know wtf defer does

	header := make([]byte, 24)  // select the first 24 bytes from the .pcap fiee
	n, err := file.Read(header) // read the first 24 bytes header selected
	if err != nil {
		panic(err) // this will usually be a read error (?)
	}

	fmt.Printf("read %d bytes\n", n)   // confirms we read only 24 bytes
	fmt.Printf("header: %x\n", header) //prints the shit in header

	fmt.Printf("BINARY:\n")
	buf := header[0:4]     //magic number
	ver := header[4:6]     //major version. 2 bytes
	mver := header[6:8]    // minor version. 2 bytes
	time := header[8:12]   // utc timezone offset -- buncha bullshit. 4 bytes
	atime := header[12:16] // timezone accuracy. 4 bytes
	ln := header[16:20]    //snapshot length. 4 bytes
	typ := header[20:24]   // link layer type. 4 bytes
	// parsed data:
	// we skip the magic number
	comb := make([]byte, 4)
	copy(comb[0:2], ver)  // we cant take from header, so
	copy(comb[2:4], mver) // we find the bytes from the
	//version := comb       // variable itself

	fmt.Printf("M-#: %x \n", buf)
	fmt.Printf("Ver:\n	Major: %x\n	Minor: %x\n", ver, mver)
	fmt.Printf("Timezone offset: %x (UTC)\n", time)
	fmt.Printf("Timezone Accuracy: %x (UTC)\n", atime)
	fmt.Printf("Snapshot length: %x\n", ln)
	fmt.Printf("Link-Layer type: %x\n", typ)

	fmt.Printf("Parsed (slightly):\n\n")
	fmt.Printf("Magic-Number [0:4]: %x\n", buf)
	fmt.Printf("Version [4:8]: %v.%v\n", binary.LittleEndian.Uint16(ver), binary.LittleEndian.Uint16(mver))
	fmt.Printf("Time offset [8:12]: %v (UTC)\nTime Accuracy [12:16]: %v (UTC)\n", parse(time), parse(atime))
	fmt.Printf("Snapshot length [16:20]: %v\n", parse(ln))
	fmt.Printf("Link-layer type [20:24]: %v\n", parse(typ))
}

// 1. Add "uint32" to the function definition so it returns a value
func parse(data []byte) uint32 {
	// 2. Use the return keyword
	return binary.LittleEndian.Uint32(data)
}
