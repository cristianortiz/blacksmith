package blacksmith

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

const version = "1.0.0"

//type Blacksmith is a struct to define parameters of Blacksmith module accesible from
//the webApp created
type Blacksmith struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux
	config   config
}

//private configurations settings for blackmist module
type config struct {
	port     string
	renderer string
}

//New() receives the working directory in filesystem of the WebApp to be created
//, defines the folder structure of it and check and .env file exists
func (bls *Blacksmith) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}
	err := bls.Init(pathConfig)
	if err != nil {
		return err
	}

	//check for .env file in the generated app working directory
	err = bls.checkDotEnv(rootPath)
	if err != nil {
		return err
	}
	//read .env
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	//create loggers and atacched to blacksmith type struct
	infoLog, errorLog := bls.startLoggers()
	bls.InfoLog = infoLog
	bls.ErrorLog = errorLog

	//Getenv return the value of DEBUG has string must be converted to a bool
	bls.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	bls.Version = version
	bls.RootPath = rootPath
	//bls.router returns a http.handler type, so we cast it to chi.mux type
	bls.Routes = bls.routes().(*chi.Mux)

	//private settings for the blacksmith module, form the .env file
	bls.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
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

//ListenAndServe start the webserver for the webapp
func (bls *Blacksmith) ListenAndServe() {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", bls.config.port),
		ErrorLog:     bls.ErrorLog,
		Handler:      bls.routes(),
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	bls.InfoLog.Printf("==> Listening on port %s", bls.config.port)
	err := server.ListenAndServe()
	if err != nil {
		bls.ErrorLog.Fatal(err)
	}

}

//checkDotenv verify if a .env files exists inside the wd of the webApp
func (bls *Blacksmith) checkDotEnv(path string) error {
	//bls method to check if a .env files exists
	err := bls.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil

}

//startLoggers creates two types of loggers
func (bls *Blacksmith) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "Error:\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}
