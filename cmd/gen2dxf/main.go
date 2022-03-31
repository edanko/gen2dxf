package main

import (
	"fmt"
	"log"

	"github.com/edanko/gen2dxf/cmd"
)

var (
	version   = "v0.0.0"
	commit    = "-"
	branch    = "main"
	buildDate = "0001-01-01T00:00:00+0000"
)

func main() {
	fmt.Println("Version:\t", version)
	fmt.Println("commit:\t", commit)
	fmt.Println("branch:\t", branch)
	fmt.Println("build date:\t", buildDate)

	root := cmd.RootCmd()
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
