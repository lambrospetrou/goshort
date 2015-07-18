package main

import (
	"flag"
	"fmt"
	"github.com/lambrospetrou/goshort/spito"
	"os"
)

func main() {

	var l = flag.String("l", "", "specify a long URL to short-spit it")
	var exp = flag.String("exp", "86400", "specify expiry time in seconds")
	flag.Parse()

	resp, err := spito.Spitit(*l, *exp, true)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while shortening with Spi.to :: \n", err.Error())
	} else {
		fmt.Fprintf(os.Stdout, "Spit-link: %s\n", resp)
	}
}
