package main

import (
	"flag"
	"github.com/Giulianos/module-packer/internal/packer"
	"log"
	"os"
)

func main() {
	specFilePath := flag.String("spec", "", "file containing packing specification")
	flag.Parse()

	if *specFilePath == "" {
		flag.Usage()
		os.Exit(1)
	}

	packingSpec, err := packer.LoadPackingSpecFromFile(*specFilePath)
	if err != nil {
		log.Println(err)
		log.Fatalln("error reading packing spec")
	}

	err = packer.Pack(*packingSpec)
	if err != nil {
		log.Println(err)
		log.Fatalln("error packing modules")
	}
}
