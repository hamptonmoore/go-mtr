package main

import (
	"fmt"
	"os"

	"github.com/hamptonmoore/go-mtr"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Missing host argument\nProper usage: /mtr [HOST]")
		os.Exit(1)
	}

	mtr_data := mtr.New(2, os.Args[1])
	<-mtr_data.Done
	for i := 0; i < len(mtr_data.Hosts); i++ {
		host := mtr_data.Hosts[i]
		fmt.Printf("%d: %s (%s) %.2fms \n", i, host.IP, host.Name, (host.Mean / 1000))
	}
}
