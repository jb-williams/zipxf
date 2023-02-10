package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func zipinArchive(workingDir *string, action_zip *string) {
	fpf(os.Stdout, "Creating zip archive... %s\n", (*workingDir + "/" + *action_zip))

	archive, err := os.Create(*workingDir + "/" + *action_zip)
	if err != nil {
		lff("archive: Create(): failed %w", err.Error())
	}
	defer archive.Close()

	fileNames, err := os.ReadDir(*workingDir)
	if err != nil {
		lff("fileNames: ReadDir(): failed %w", err.Error())
	}

	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	for _, file := range fileNames {
		fpf(os.Stdout, "Archiving file.... %s\n", file)

		files, err := os.Open(file.Name())
		if err != nil {
			lff("files: Open(): failed %w", err.Error())
		}
		defer files.Close()

		myWriter, err := zipWriter.Create(files.Name())
		if err != nil {
			lff("myWriter: Create(): failed %w", err.Error())
		}

		if _, err := io.Copy(myWriter, files); err != nil {
			lff("err: Copy(myWriter, files): failed %w", err.Error())
		}
	}
}

func unZipArchive(workingDir *string, action_unzip *string) {
	fpf(os.Stdout, "Opening zip archive... %s\n", *action_unzip)

	dest := *action_unzip
	archive, err := zip.OpenReader(*workingDir + "/" + *action_unzip)
	if err != nil {
		lff("archive: OpenReader(): failed %w", err.Error())
	}
	defer archive.Close()

	for _, file := range archive.File {
		filePath := filepath.Join(strings.TrimSuffix(dest, filepath.Ext(dest)), file.Name)
		fpf(os.Stdout, "Unzipping file... %s\n", filePath)

		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				lff("fileinfo.isdir: MkdirAll(): failed %w", err.Error())
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			lff("mkdirall: filepath.dir(): failed %w", err.Error())
		}

		destFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			lff("destFile: OpenFile(): failed %w", err.Error())
		}

		fileInArchive, err := file.Open()
		if err != nil {
			lff("fileInArchive: Getwd(): failed %w", err.Error())
		}

		if _, err := io.Copy(destFile, fileInArchive); err != nil {
			lff("writingout: Copy(): failed %w", err.Error())
		}

		destFile.Close()
		fileInArchive.Close()
	}
}
