package blacksmith

import "os"

//CreateDirIfNotExists checks and create any directory of the desired webApp folder structure
func (bls *Blacksmith) CreateDirIfNotExist(path string) error {
	const mode = 0755
	//returns the folder info
	_, err := os.Stat(path)
	//if the path does not exists (the error  returned by os.Stat)
	if os.IsNotExist(err) {
		//create the folder
		err = os.Mkdir(path, mode)
		if err != nil {
			return err
		}
	}
	return nil
}

//CreateFileIfNotExists checks if the filename given as parameter exists and if is'nt, is created it
func (bls *Blacksmith) CreateFileIfNotExists(path string) error {
	//get the info of the filename
	var _, err = os.Stat(path)
	//if an error is returned means the file does'nt exists
	if os.IsNotExist(err) {
		//create the file, its returns a *os.File type
		var file, err = os.Create(path)
		if err != nil {
			return err
		}
		//close the file when the funtion work ends
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}

	return nil
}
