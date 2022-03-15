package main

import (
	"embed"
	"io/ioutil"
)

//go:embed templates
var templateFS embed.FS

//copyFilefromTemplate() extract the content of a file
func copyFilefromTemplate(templatePath, targetFile string) error {
	//TODO : check to ensure files does not already exist

	data, err := templateFS.ReadFile(templatePath)
	if err != nil {
		exitGracefully(err)
	}
	//copy the data to a file
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
