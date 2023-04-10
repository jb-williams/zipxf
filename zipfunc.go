package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func zipinArchive(zipWriter *zip.Writer, workingDir, myArchive, osSep string) {

	fileNames, err := os.ReadDir(workingDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range fileNames {
		if !file.IsDir() && file.Name() != myArchive && file.Name() != ".git" {
			//fpf(os.Stdout, "Archiving file.... %s\n", file)

			files, err := os.ReadFile(workingDir + osSep + file.Name())
			if err != nil {
				log.Fatal(err)
			}

			writer, err := zipWriter.Create(file.Name())
			if err != nil {
				log.Fatal(err)
			}
			_, err = writer.Write(files)
			if err != nil {
				log.Fatal(err)
			}
			// if _, err := io.Copy(writer, files); err != nil {
			// 	log.Fatal(err)
			// }
		} else if file.IsDir() && file.Name() != myArchive && file.Name() != ".git" {
			newBase := workingDir + osSep + file.Name()
			if err := os.MkdirAll(newBase, os.ModePerm); err != nil {
				log.Fatal(err)
			}

			zipinArchive(zipWriter, newBase, newBase, osSep)
			//zipinArchive(zipWriter, newBase, myArchive, osSep)
		}
	}
}

//func zipinArchive(workingDir *string, action_zip *string, osSep string) {
//fpf(os.Stdout, "Creating zip archive... %s\n", *action_zip)

//fileNames, err := os.ReadDir(*workingDir)
//if err != nil {
//log.Fatal(err)
//}
//
//archive, err := os.Create(*workingDir + osSep + *action_zip)
//if err != nil {
//log.Fatal(err)
//}
//defer archive.Close()
//
//zipWriter := zip.NewWriter(archive)
//defer zipWriter.Close()

//for _, file := range fileNames {
//fpf(os.Stdout, "Archiving file.... %s\n", file)

//files, err := os.Open(file.Name())
//if err != nil {
//log.Fatal(err)
//}
//defer files.Close()
//
//writer, err := zipWriter.Create(files.Name())
//if err != nil {
//log.Fatal(err)
//}

//if _, err := io.Copy(writer, files); err != nil {
//log.Fatal(err)
//}
//}
//}

func unZipArchive(workingDir *string, action_unzip *string, osSep string) {
	dest := *action_unzip

	//fpf(os.Stdout, "Opening zip archive... %s\n", *action_unzip)

	archive, err := zip.OpenReader(*workingDir + osSep + *action_unzip)
	if err != nil {
		log.Fatal(err)
	}
	defer archive.Close()

	for _, file := range archive.File {
		filePath := filepath.Join(strings.TrimSuffix(dest, filepath.Ext(dest)), file.Name)

		//fpf(os.Stdout, "Unzipping file... %s\n", filePath)

		if file.FileInfo().IsDir() {
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
