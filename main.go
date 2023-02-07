package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// var pl = fmt.Println
var fpf = fmt.Fprintf

func main() {
	help := flag.Bool("h", false, "Show help")
	action_zip := flag.String("z", "archive.zip", "Zip files into and Archive.")
	action_unzip := flag.String("u", "", "Unzip archive.")
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

func ZipinArchive(workingDir *string, action_zip *string) {
	archive, err := os.Create(*workingDir + "/" + *action_zip)
	fpf(os.Stdout, "Creating zip archive... %s\n", (*workingDir + "/" + *action_zip))
	if err != nil {
		log.Fatal(err)
	}
	defer archive.Close()

	fileNames, err := os.ReadDir(*workingDir)
	if err != nil {
		log.Fatal(err)
	}

	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	for _, file := range fileNames {
		fpf(os.Stdout, "Archiving file.... %s\n", file)
		files, err := os.Open(file.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer files.Close()

		writer, err := zipWriter.Create(files.Name())
		if err != nil {
			log.Fatal(err)
		}

		if _, err := io.Copy(writer, files); err != nil {
			log.Fatal(err)
		}
	}
}

func UnZipArchive(workingDir *string, action_unzip *string) {
	dest := *action_unzip
	fpf(os.Stdout, "Opening zip archive... %s\n", *action_unzip)
	//fpf(os.Stdout, "Opening zip archive... %s\n", (workingDir + "/" + *outputArchive))
	archive, err := zip.OpenReader(*workingDir + "/" + *action_unzip)
	if err != nil {
		log.Fatal(err)
	}
	defer archive.Close()

	for _, file := range archive.File {
		filePath := filepath.Join(strings.TrimSuffix(dest, filepath.Ext(dest)), file.Name)
		//pl("unzipping file ", filePath)
		fpf(os.Stdout, "Unzipping file... %s\n", filePath)

		if file.FileInfo().IsDir() {
			//pl("creating directory...")
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				log.Fatal(err)
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			log.Fatal(err)
		}

		destFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			log.Fatal(err)
		}

		fileInArchive, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}

		if _, err := io.Copy(destFile, fileInArchive); err != nil {
			log.Fatal(err)
		}

		destFile.Close()
		fileInArchive.Close()
	}
}
