package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var fpf = fmt.Fprintf
var lff = log.Fatalf

func main() {
	help := flag.Bool("h", false, "Show help")
	action_zip := flag.String("z", "", "(redundant zipping)Zip all files in current working dir into an Archive.")
	action_unzip := flag.String("uz", "", "Unzip Archive into current working dir.")
	action_targz := flag.String("t", "", "(redundant targzing)Tar Gz all in current working dir.")
	action_utargz := flag.String("ut", "", "Un-Tar Gz all in current working dir.")
	flag.Parse()

	workingDir, err := os.Getwd()
	if err != nil {
		lff("workingDir: Getwd(): failed %w", err.Error())
	}

	switch {
	case *action_zip != "":
		zipinArchive(&workingDir, action_zip)

	case *action_unzip != "":
		unZipArchive(&workingDir, action_unzip)

	case *action_targz != "":
		tarGzipArchive(&workingDir, action_targz)

	case *action_utargz != "":
		utarGzipArchive(&workingDir, action_utargz)

	case *help:
		flag.Usage()
		os.Exit(0)

	default:
		flag.Usage()
		os.Exit(0)

	}
}
