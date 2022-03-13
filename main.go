package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// monitor for CTRL-C or CTRL-Z signal and exit if received
	fmt.Println("CTRL-C or CTRL-Z to exit")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTSTP, syscall.SIGINT)
	// loop endlessly
	for {
		// process some things
		select {
		case <-sig:
			fmt.Println("Exiting...")
			os.Exit(0)
		}
	}
}
