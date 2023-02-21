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

func tarGzipArchive(workingDir *string, action_targz *string, osSep string) {
	fpf(os.Stdout, "Creating zip archive... %s\n", *action_targz)
	fileNames, err := os.ReadDir(*workingDir)
	if err != nil {
		lff("Tar: fileNames: ReadDir(): failed: %w", err.Error())
	}

	archive, err := os.Create(*workingDir + osSep + *action_targz)
	if err != nil {
		lff("Tar: archive: Create(): failed: %w", err.Error())
	}
	defer archive.Close()

	gzWriter := gzip.NewWriter(archive)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	for _, file := range fileNames {
		fpf(os.Stdout, "Archiving file.... %s\n", file)
		//if file.IsDir() {
		//}
		files, err := os.Open(file.Name())
		if err != nil {
			lff("Tar: files: Open(): failed: %w", err.Error())
		}
		defer files.Close()
		info, err := files.Stat()
		if err != nil {
			lff("Tar: info: files.Stat(): failed: %w", err.Error())
		}

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			lff("Tar: header: tar.FileInfoHeader(): failed: %w", err.Error())
		}
		header.Name = file.Name()

		err = tarWriter.WriteHeader(header)
		if err != nil {
			lff("Tar: tarWriter: WriteHeader(): failed: %w", err.Error())
		}

		if _, err := io.Copy(tarWriter, files); err != nil {
			log.Fatal(err)
			lff("Tar: Last: Copy(): failed: %w", err.Error())
		}
	}

}

func utarGzipArchive(workingDir *string, action_utargz *string, osSep string) {
	fpf(os.Stdout, "Un-Tar.Gz-ing... %s\n", (*action_utargz))
	myReader, err := os.Open(*action_utargz)
	if err != nil {
		lff("Untar: myReader: Open(): failed: %w", err.Error())
	}
	ugziper, err := gzip.NewReader(myReader)
	if err != nil {
		lff("Untar: ugziper: NewReader(): failed %w", err.Error())
	}
	defer ugziper.Close()

	untargz := tar.NewReader(ugziper)
	if err != nil {
		lff("Untar: untargz: NewReader(): failed %w", err.Error())
	}
	for {
		header, err := untargz.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			lff("Untar: header: untargz.Next(): failed: %w", err.Error())
		}
		filePath := filepath.Join(strings.TrimSuffix(*action_utargz, filepath.Ext(*action_utargz)), header.Name)
		if header.FileInfo().IsDir() {
			//pl("creating directory...")
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				log.Fatal(err)
				lff("Untar: header.Fileinfo: MkdirAll(): failed: %w", err.Error())
			}
			continue
		}
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			log.Fatal(err)
			lff("Untar: filepath.Dir(filePath): MkdirAll(): failed: %w", err.Error())
		}
		destFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, header.FileInfo().Mode())
		if err != nil {
			log.Fatal(err)
			lff("Untar: destFile: OpenFile(): failed: %w", err.Error())
		}

		if _, err := io.Copy(destFile, untargz); err != nil {
			lff("Untar: Last: Copy(): failed: %w", err.Error())
		}

		destFile.Close()
	}
}
