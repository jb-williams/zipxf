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

	archive, err := os.Create(*workingDir + "/" + *action_targz)
	if err != nil {
		lff("archive: Create(): failed %w", err.Error())
	}
	defer archive.Close()

	fileNames, err := os.ReadDir(*workingDir)
	if err != nil {
		log.Fatal(err)
		lff("fileNames: ReadDir(): failed %w", err.Error())
	}

	gzWriter := gzip.NewWriter(archive)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	for _, file := range fileNames {
		fpf(os.Stdout, "Archiving file.... %s\n", file)

		files, err := os.Open(file.Name())
		if err != nil {
			lff("files: Open(): failed %w", err.Error())
		}
		defer files.Close()
		info, err := files.Stat()
		if err != nil {
			lff("info: files.Stat(): failed %w", err.Error())
		}

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			lff("header: NewReader(): failed %w", err.Error())
		}
		header.Name = file.Name()

		err = tarWriter.WriteHeader(header)
		if err != nil {
			lff("tarWriter: WriteHeater(): failed %w", err.Error())
		}

		if _, err := io.Copy(tarWriter, files); err != nil {
			lff("tarWriter: Copy(): failed %w", err.Error())
		}
	}

}

func utarGzipArchive(workingDir *string, action_utargz *string) {
	fpf(os.Stdout, "Un-Tar.Gz-ing... %s\n", (*workingDir + "/" + *action_utargz))

	myReader, err := os.Open(*action_utargz)
	if err != nil {
		lff("myReader: Open(): failed %w", err.Error())
	}

	ugziper, err := gzip.NewReader(myReader)
	if err != nil {
		lff("ugziper: NewReader(): failed %w", err.Error())
	}
	defer ugziper.Close()

	untargz := tar.NewReader(ugziper)
	if err != nil {
		lff("untargz: NewReader(): failed %w", err.Error())
	}

	for {
		header, err := untargz.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			lff("header: err != nil %w", err.Error())
		}

		filePath := filepath.Join(strings.TrimSuffix(*action_utargz, filepath.Ext(*action_utargz)), header.Name)

		if header.FileInfo().IsDir() {
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				lff("headerfileinfoisdir: MkdirAll() %w", err.Error())
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			lff("headerfileinfoisdir: filepath.Dir() %w", err.Error())
		}

		destFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, header.FileInfo().Mode())
		if err != nil {
			lff("destFile: OpenFile() %w", err.Error())
		}

		if _, err := io.Copy(destFile, untargz); err != nil {
			lff("Last: Copy(): failed %w", err.Error())
		}

		destFile.Close()
	}
}
