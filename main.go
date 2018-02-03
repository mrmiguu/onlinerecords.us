package main

import (
	"github.com/mrmiguu/jsutil"
	"github.com/mrmiguu/sock"
)

func main() {
	if err := jsutil.CompileWithGzip(sock.Root + "/script.go"); err != nil {
		panic(err)
	}

	sock.Addr = ":80"
	handshake := sock.Rbool()
	for range handshake {
		println("handshake")
	}
}
