package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var fpf = fmt.Fprintf

func main() {
	help := flag.Bool("h", false, "Show help")
	action_zip := flag.String("zip", "archive.zip", "Zip all files in current working dir into an Archive.\n\tDefault: archive.zip")
	action_unzip := flag.String("unzip", "", "Unzip Archive into current working dir.")
	flag.Parse()

	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case *action_zip != "":
		ZipinArchive(&workingDir, action_zip)

	case *action_unzip != "":
		UnZipArchive(&workingDir, action_unzip)

	case *help:
		flag.Usage()
		os.Exit(0)

	default:
		flag.Usage()
		os.Exit(0)

	}
}
