package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/edanko/gen2dxf/pkg/gen"
)

func main() {

	var wcogs []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		if info.Mode().IsRegular() {
			if filepath.Base(path)[:5] == "wcog1" && filepath.Ext(path) == ".csv" {
				wcogs = append(wcogs, path)
			}
		}

		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
	if len(wcogs) == 0 {
		fmt.Println("[i] no wcogs found!")
	}

	var gens []string
	err = filepath.Walk(".", func(path string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		if info.Mode().IsRegular() {
			if filepath.Ext(path) == ".gen" {
				gens = append(gens, path)
			}
		}

		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
	if len(gens) == 0 {
		fmt.Println("[i] no gens found!")
		os.Exit(0)
	}

	//wcog := wcog.ReadWCOGs(wcogs)

	l := len(gens)

	for i, p := range gens {

		fmt.Printf("[*] %d/%d...\n", i+1, l)

		g := gen.ParseProfileFile(p)

		_ = g
		//err = draw.PlateToDXF(g, wcog)
		//if err != nil {
		//	log.Fatalln(err)
		//}

	}

}
