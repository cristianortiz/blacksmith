package blacksmith

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cristianortiz/blacksmith/render"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

const version = "1.0.0"

//type Blacksmith is a struct to define parameters of Blacksmith module accesible later
//when the webapp "myapp" calls blacksmith module
type Blacksmith struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux //http.handler (chi mux), to config and init a webserver
	Render   *render.Render
	config   config
}

//private configurations settings for blackmist module
type config struct {
	port     string
	renderer string
}

//-New() receives the working directory in filesystem of the WebApp to be created
//defines the folder structure of webapp,check if .env file exists in wd,creates
//and init webapp loggers, set the handler (*chi mux type) to config webserver to tun the webapp
func (bls *Blacksmith) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}
	//create the folder structure for the webapp
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
	//Getenv return the value of DEBUG as string, must be converted to a bool
	bls.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	bls.Version = version
	bls.RootPath = rootPath
	//bls.router returns a http.handler type, must be cast it to chi.mux type to atach
	//to blacksmith object as a handler for the webapp webserver
	bls.Routes = bls.routes().(*chi.Mux)
	//private settings for the blacksmith module, from the .env file
	bls.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}

	// also this more simple way bls.createRenderer() using the alternative version in line 147
	bls.Render = bls.createRenderer(bls)
	return nil
}

//Init loops to folderNames property of initPath struct and creates the webApp folder structure
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

//ListenAndServe config and start the webserver using Blacksmith.Routes as handler
// (a *chi mux type in this case )for running the webapp
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
	//bls method to check if a .env files exists, or created it
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

//createRenderer config anc creates a Render struct to atach to blacksmith struct
func (bls *Blacksmith) createRenderer(b *Blacksmith) *render.Render {
	myRenderer := render.Render{
		Renderer: b.config.renderer,
		RootPath: b.RootPath,
		Port:     b.config.port,
	}
	return &myRenderer
}

//this is an alternative way to above function, using only the receivers params to create
//the attach the render struct into blacksmith struct
// func (b *blacksmith) createRenderer() {
// 	b.Render = &render.Render{
// 		   Renderer: b.config.renderer,
// 		   RootPath: b.RootPath,
// 		   Port: b.config.port,
// 	}
// }
