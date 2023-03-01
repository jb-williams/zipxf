package main

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"flag"
	"fmt"
	"log"
	"os"
)

var fpf = fmt.Fprintf
var lff = log.Fatalf

func main() {
	help := flag.Bool("h", false, "Show help")
	action_zip := flag.String("z", "", "Zip all files in current working dir into an Archive.")
	action_unzip := flag.String("uz", "", "Unzip Archive into current working dir.")
	action_targz := flag.String("t", "", "Tar Gz all in current working dir.")
	action_utargz := flag.String("ut", "", "Un-Tar Gz all in current working dir.")
	flag.Parse()

	workingDir, err := os.Getwd()
	if err != nil {
		lff("workingDir: Getwd(): failed %w", err.Error())
	}

	osSep := string(os.PathSeparator)

	switch {
	case *action_zip != "":
		archive, err := os.Create(workingDir + osSep + *action_zip)
		if err != nil {
			log.Fatal(err)
		}
		defer archive.Close()

		zipWriter := zip.NewWriter(archive)
		defer zipWriter.Close()

		// zipinArchive(zipWriter, workingDir, "", myArchive, osSep)
		zipinArchive(zipWriter, workingDir, *action_zip, osSep)

		if err != nil {
			log.Println(err)
		}

		err = zipWriter.Close()
		if err != nil {
			log.Println(err)
		}
		//zipinArchive(&workingDir, action_zip, osSep)

	case *action_unzip != "":
		unZipArchive(&workingDir, action_unzip, osSep)

	case *action_targz != "":
		archive, err := os.Create(workingDir + osSep + *action_targz)
		if err != nil {
			lff("Tar: archive: Create(): failed: %w", err.Error())
		}
		defer archive.Close()

		gzWriter := gzip.NewWriter(archive)
		defer gzWriter.Close()

		tarWriter := tar.NewWriter(gzWriter)
		defer tarWriter.Close()

		tarGzipArchive(tarWriter, workingDir, *action_targz, osSep)
		//		tarGzipArchive(&workingDir, action_targz, osSep)

	case *action_utargz != "":
		utarGzipArchive(&workingDir, action_utargz, osSep)

	case *help:
		flag.Usage()
		os.Exit(0)

	default:
		flag.Usage()
		os.Exit(0)

	}
}
