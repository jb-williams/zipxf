package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
)

func tarGzipArchive(workingDir *string, action_targz *string) {
<<<<<<< HEAD
	fpf(os.Stdout, "Creating zip archive... %s\n", (*workingDir + "/" + *action_targz))
=======
	fpf(os.Stdout, "Creating tar.gz archive... %s\n", (*workingDir + "/" + *action_targz))
>>>>>>> targzip
	archive, err := os.Create(*workingDir + "/" + *action_targz)
	if err != nil {
		log.Fatal(err)
	}
	defer archive.Close()

	fileNames, err := os.ReadDir(*workingDir)
	if err != nil {
		log.Fatal(err)
	}
<<<<<<< HEAD
	//var buf bytes.Buffer
	gzWriter := gzip.NewWriter(archive)
	defer gzWriter.Close()
	tarWriter := tar.NewWriter(gzWriter)
	//tarWriter := tar.NewWriter(archive)
=======

	gzWriter := gzip.NewWriter(archive)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
>>>>>>> targzip
	defer tarWriter.Close()

	for _, file := range fileNames {
		fpf(os.Stdout, "Archiving file.... %s\n", file)
		files, err := os.Open(file.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer files.Close()
<<<<<<< HEAD
=======

>>>>>>> targzip
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

<<<<<<< HEAD
func utarGzipArchive(workingDir *string, action_unzip *string) {

=======
func utarGzipArchive(workingDir *string, action_utargz *string, myReader io.Reader) {
	//func utarGzipArchive(workingDir *string, action_utargz *string, myReader io.Reader) {
	//fpf(os.Stdout, "Un-Tar.Gz-ing... %s\n", (*workingDir + "/" + *action_utargz))
	// archive, err := os.OpenFile((*workingDir + "/" + *action_utargz), os.O_RDWR|os.O_TRUNC, 0777)
	// if err != nil {
	// 	log.Fatal("archive fail")
	// }

	ugziper, err := gzip.NewReader(myReader)
	if err != nil {
		log.Fatal("newreader failed")
	}
	defer ugziper.Close()

	untargz := tar.NewReader(ugziper)

	for {
		header, err := untargz.Next()

		// if err != nil {
		// 	log.Fatal("EOF fail")
		// }
		if err == io.EOF {
			log.Fatal("EOF fail")
			//log.Fatal(err)
		}

		if err != nil {
			lff("unTarGzip: Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(header.Name, 0755); err != nil {
				//log.Fatal(err)
				lff("unTarGzip: Mkdir() failed: %s", err.Error())
			}

		case tar.TypeReg:
			outFile, err := os.Create(header.Name)
			if err != nil {
				//log.Fatal(err)
				lff("unTarGzip: Create() failed: %s", err.Error())
			}
			if _, err := io.Copy(outFile, untargz); err != nil {
				//log.Fatal(err)
				lff("unTarGzip: Copy() failed: %s", err.Error())
			}
			outFile.Close()
		default:
			lff("unTarGzip: unknown type: %b in %s", header.Typeflag, header.Name)
		}
	}
>>>>>>> targzip
}
