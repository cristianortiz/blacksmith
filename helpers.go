package blacksmith

import "os"

//CreateDirIfNotExists create any missing directory of the desired webApp structure
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
