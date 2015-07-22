package main

import (
	"flag"
	"fmt"
	"github.com/lambrospetrou/goshort/spito"
	"os"
)

func main() {

	var v = flag.String("v", "", "specify the Spit ID you want to see details")
	var c = flag.String("c", "", "specify a long URL (or text) to short-spit it")
	var e = flag.Uint64("e", 86400, "specify expiry time in seconds")
	var t = flag.String("t", "url", "specify the type of the Spit (url or text)")
	flag.Parse()

	if len(*v) > 0 {
		resp, err := spito.View(*v)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while fetching information for: %s\n\t%s", v, err.Error())
		} else {
			fmt.Fprintf(os.Stdout, "Spit-details: %s\n", resp)
		}
	} else {
		if len(*c) == 0 {
			fmt.Fprintf(os.Stderr, "Cannot create empty Spit!\n")
			return
		}
		resp, err := spito.Spit(*c, *t, *e, true)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error while shortening with Spi.to :: \n", err.Error())
		} else {
			fmt.Fprintf(os.Stdout, "Spit-link: %s\n", resp)
		}
	}
}
