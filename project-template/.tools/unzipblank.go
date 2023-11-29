package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

const targetOutput string = "./.."

func main() {

	var fZip *zip.ReadCloser
	var err error

	fZip, err = zip.OpenReader("./blank.zip")

	if err != nil {
		panic(fmt.Sprintf("PANIC!!! * Runs in circles * => %v", err.Error()))
	}

	defer fZip.Close()

	for _, file := range fZip.Reader.File {

		if file.FileInfo().IsDir() {
			fmt.Printf("Dir\t'%v' ", file.Name)
			var dirPath = filepath.Join(targetOutput, file.Name)
			err := os.MkdirAll(dirPath, file.FileInfo().Mode())
			if err != nil {
				panic(fmt.Sprintf("could not create directory %v => %v", dirPath, err.Error()))
			}
			fmt.Println("=> created / exists")
			continue

		} else {
			fmt.Printf("File\t'%v' ", file.Name)

			// Open Input File (from Zip)
			var fFile fs.File
			fFile, err = fZip.Open(file.Name)

			if err != nil {
				fmt.Println("=> Failed to open => ", err.Error())
				continue
			}

			defer fFile.Close()

			// Open/Create Output File (On Harddrive)
			var filePath = filepath.Join(targetOutput, file.Name)
			var fOut, err = os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, file.FileInfo().Mode())

			if err != nil {
				panic(fmt.Sprintf("=> could not creat outputfile %v => %v", filePath, err.Error()))
			}

			defer fOut.Close()

			// Copy Data
			_, err = io.Copy(fOut, fFile)
			if err != nil {
				fmt.Printf("=> failed to copy data => %v\n", err.Error())
				//TODO: cleanup created files here
				continue
			}
			fmt.Println("=> done ")
		}
	}
}
