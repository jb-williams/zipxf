package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func tarGzipArchive(workingDir *string, action_targz *string) {
	fpf(os.Stdout, "Creating zip archive... %s\n", (*workingDir + "/" + *action_targz))
	fileNames, err := os.ReadDir(*workingDir)
	if err != nil {
		log.Fatal(err)
	}

	archive, err := os.Create(*workingDir + "/" + *action_targz)
	if err != nil {
		log.Fatal(err)
	}
	defer archive.Close()

	gzWriter := gzip.NewWriter(archive)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
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

func utarGzipArchive(workingDir *string, action_utargz *string) {
	fpf(os.Stdout, "Un-Tar.Gz-ing... %s\n", (*workingDir + "/" + *action_utargz))
	//myReader, err := tar.OpenReader(*action_utargz)
	myReader, err := os.Open(*action_utargz)
	if err != nil {
		log.Fatal("openreader failed")
	}
	ugziper, err := gzip.NewReader(myReader)
	// ugziper, err := gzip.NewReader(io.Reader(*action_utargz))
	if err != nil {
		lff("ugziper: NewReader(): failed %w", err.Error())
	}
	defer ugziper.Close()

	untargz := tar.NewReader(ugziper)
	if err != nil {
		lff("untargz: NewReader(): failed %w", err.Error())
	}
	// Open and iterate through the files in the archive.
	for {
		header, err := untargz.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			lff("err not nil %w", err.Error())
		}
		filePath := filepath.Join(strings.TrimSuffix(*action_utargz, filepath.Ext(*action_utargz)), header.Name)
		if header.FileInfo().IsDir() {
			//pl("creating directory...")
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				log.Fatal(err)
			}
			continue
		}
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			log.Fatal(err)
		}
		destFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, header.FileInfo().Mode())
		if err != nil {
			log.Fatal(err)
		}

		if _, err := io.Copy(destFile, untargz); err != nil {
			lff("Last: Copy(): failed %w", err.Error())
		}

		destFile.Close()
	}
}
