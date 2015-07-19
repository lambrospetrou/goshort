package main

import (
	"flag"
	"fmt"
	"github.com/lambrospetrou/goshort/spito"
	"os"
)

func main() {

	var l = flag.String("c", "", "specify a long URL (or text) to short-spit it")
	var e = flag.String("e", "86400", "specify expiry time in seconds")
	var t = flag.String("t", "url", "specify the type of the Spit (url or text)")
	flag.Parse()

	resp, err := spito.Spitit(*l, *t, *e, true)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while shortening with Spi.to :: \n", err.Error())
	} else {
		fmt.Fprintf(os.Stdout, "Spit-link: %s\n", resp)
	}
}
