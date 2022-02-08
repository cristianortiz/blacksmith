package blacksmith

//initPath is a struct to define the main path and folders structure for the WebApp
type initPaths struct {
	//the working directory path to create the webApp folder structure
	rootPath    string
	folderNames []string
}

type cookieConfig struct {
	name     string
	lifetime string
	persist  string
	secure   string
	domain   string
}
