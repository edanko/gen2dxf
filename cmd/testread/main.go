package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/edanko/gen2dxf/pkg/draw"
	"github.com/edanko/gen2dxf/pkg/gen"
	"github.com/edanko/gen2dxf/pkg/wcog"
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

	wcog, err := wcog.ReadWCOGs(wcogs)
	if err != nil {
		log.Fatalln(err)
	}

	l := len(gens)

	for i, fname := range gens {
		fmt.Printf("[i] %d/%d (%s)\n", i+1, l, filepath.Base(fname))

		g := gen.ParsePlateFile(fname)

		if g == nil {
			fmt.Println(" - skipping")
			continue
		}

		err = draw.PlateToDXF(g, wcog)
		if err != nil {
			log.Fatalln(err)
		}

	}
}
