package main

import (
	"github.com/mrmiguu/page"
	"github.com/mrmiguu/page/css"
	"github.com/mrmiguu/sock"
)

var (
	bg    = page.Class("bg")
	logo  = page.Class("logo")
	strip = page.Class("strip")
)

func main() {
	sock.Addr = "47.148.135.216"
	handshake := sock.Wbool()
	handshake <- true
	for {
		select {
		case <-logo.Hit:
			go logo.Display(css.None)
			bg.Play()
		}
	}
}
