package main

import (
	"embed"
	"errors"
	"io/ioutil"
	"os"
)

//go:embed templates
var templateFS embed.FS

//copyFilefromTemplate() extract the content from a template file and
//insert it into another file
func copyFilefromTemplate(templatePath, targetFile string) error {

	if fileExists(targetFile) {
		return errors.New(targetFile + " already exists")
	}

	data, err := templateFS.ReadFile(templatePath)
	if err != nil {
		exitGracefully(err)
	}
	//copy the data into a new a file
	err = copyDataToFile(data, targetFile)
	if err != nil {
		exitGracefully(err)
	}

	return nil
}

//copyDataToFile() copy the data extracted from a file as []byte in another destinations file
func copyDataToFile(data []byte, to string) error {
	err := ioutil.WriteFile(to, data, 0664)
	if err != nil {
		return err
	}
	return nil
}

//fileExists checks if a given file path exists
func fileExists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true

}
