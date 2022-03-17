package main

import (
	"embed"
	"io/ioutil"
	"os"
)

//go:embed templates
var templateFS embed.FS

//copyFilefromTemplate() extract the content from a template file and
//insert it into another file
func copyFilefromTemplate(templatePath, targetFile string) error {
	//TODO : check to ensure files does not already exist
	var _, err = os.Stat(templatePath)
	//if an error is returned means the already exists
	if os.IsNotExist(err) {

		data, err := templateFS.ReadFile(templatePath)
		if err != nil {
			exitGracefully(err)
		}
		//copy the data into a new a file
		err = copyDataToFile(data, targetFile)
		if err != nil {
			exitGracefully(err)
		}

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
