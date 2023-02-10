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
	myReader, err := os.Open(*action_utargz)
	if err != nil {
		log.Fatal("openreader failed")
	}
	ugziper, err := gzip.NewReader(myReader)
	if err != nil {
		log.Fatal("new gz reader failed")
	}
	defer ugziper.Close()

	untargz := tar.NewReader(ugziper)
	if err != nil {
		log.Fatal("new tar reader failed")
	}
	var header *tar.Header
	for header, err = untargz.Next(); err == nil; header, err = untargz.Next() {
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(header.Name, 0755); err != nil {
				lff("ExtractTarGz: Mkdir() failed: %w", err.Error())
			}
		case tar.TypeReg:
			outFile, err := os.Create(header.Name)
			if err != nil {
				lff("ExtractTarGz: Create() failed: %w", err.Error())
			}

			if _, err := io.Copy(outFile, untargz); err != nil {
				// outFile.Close error omitted as Copy error is more interesting at this point
				outFile.Close()
				lff("ExtractTarGz: Copy() failed: %w", err.Error())
			}
			if err := outFile.Close(); err != nil {
				lff("ExtractTarGz: Close() failed: %w", err.Error())
			}
		default:
			lff("ExtractTarGz: uknown type: %b in %s", header.Typeflag, header.Name)
		}
	}
	if err != io.EOF {
		lff("ExtractTarGz: Next() failed: %w", err)
	}
	// //////// this goes back to untargz declaration if works
	// for {
	// 	header, err := untargz.Next()

	// 	// if err != nil {
	// 	// 	log.Fatal("EOF fail")
	// 	// }
	// 	if err == io.EOF {
	// 		log.Fatal("EOF fail")
	// 		//log.Fatal(err)
	// 	}

	// 	if err != nil {
	// 		lff("unTarGzip: Next() failed: %s", err.Error())
	// 	}

	// 	switch header.Typeflag {
	// 	case tar.TypeDir:
	// 		if err := os.Mkdir(header.Name, 0755); err != nil {
	// 			//log.Fatal(err)
	// 			lff("unTarGzip: Mkdir() failed: %s", err.Error())
	// 		}

	// 	case tar.TypeReg:
	// 		outFile, err := os.Create(header.Name)
	// 		if err != nil {
	// 			//log.Fatal(err)
	// 			lff("unTarGzip: Create() failed: %s", err.Error())
	// 		}
	// 		if _, err := io.Copy(outFile, untargz); err != nil {
	// 			//log.Fatal(err)
	// 			lff("unTarGzip: Copy() failed: %s", err.Error())
	// 		}
	// 		outFile.Close()
	// 	default:
	// 		lff("unTarGzip: unknown type: %b in %s", header.Typeflag, header.Name)
	// 	}
	//}
}
