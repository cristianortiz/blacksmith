package blacksmith

const version = "1.0.0"

//type Blacksmith is a struct to define  basics parameters of Blacksmith module
type Blacksmith struct {
	AppName string
	Debug   bool
	Version string
}

//func New() receives the path in filesystem of the WebApp to be created
//and defines the folder structure of it
func (bls *Blacksmith) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}
	err := bls.Init(pathConfig)
	if err != nil {
		return err
	}

	return nil
}

//Init loops to folderNames to initPath struct and create the webApp folder structure
func (bls *Blacksmith) Init(p initPaths) error {
	root := p.rootPath

	//create a folder if it doesn't exists
	for _, path := range p.folderNames {
		err := bls.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}

	}
	return nil
}
