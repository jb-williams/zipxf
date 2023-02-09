package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
)

func tarGzipArchive(workingDir *string, action_targz *string) {
	fpf(os.Stdout, "Creating zip archive... %s\n", (*workingDir + "/" + *action_targz))
	archive, err := os.Create(*workingDir + "/" + *action_targz)
	if err != nil {
		log.Fatal(err)
	}
	defer archive.Close()

	fileNames, err := os.ReadDir(*workingDir)
	if err != nil {
		log.Fatal(err)
	}
	//var buf bytes.Buffer
	gzWriter := gzip.NewWriter(archive)
	defer gzWriter.Close()
	tarWriter := tar.NewWriter(gzWriter)
	//tarWriter := tar.NewWriter(archive)
	defer tarWriter.Close()

	for _, file := range fileNames {
		fpf(os.Stdout, "Archiving file.... %s\n", file)
		files, err := os.Open(file.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer files.Close()
		info, err := files.Stat()
		if err != nil {
			log.Fatal(err)
		}

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			log.Fatal(err)
		}
		header.Name = file.Name()

		err = tarWriter.WriteHeader(header)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := io.Copy(tarWriter, files); err != nil {
			log.Fatal(err)
		}
	}

}

func utarGzipArchive(workingDir *string, action_unzip *string) {

}
