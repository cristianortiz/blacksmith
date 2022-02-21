package blacksmith

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
	"github.com/cristianortiz/blacksmith/render"
	"github.com/cristianortiz/blacksmith/session"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

const version = "1.0.0"

//type Blacksmith is a struct to define parameters of Blacksmith module accesible later
//for the webapp "myapp"  blacksmith module
type Blacksmith struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux //http.handler (chi mux), to config and init a webserver
	Render   *render.Render
	Session  *scs.SessionManager //SessionManage type
	DB       Database            //to database type and pool connection to myapp
	JetViews *jet.Set
	config   config
}

//private configurations  for blacksmisth module
type config struct {
	port        string
	renderer    string
	cookie      cookieConfig
	sessionType string
	database    databaseConfig
}

//New() receives the working directory in filesystem of the WebApp to be created
//defines the folder structure of webapp,check if .env file exists in wd,creates
//and init webapp loggers, set the handler (*chi mux type) to config webserver to tun the webapp
func (bls *Blacksmith) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}
	//create the folder structure for myapp
	err := bls.Init(pathConfig)
	if err != nil {
		return err
	}
	//check for .env file in the generated app working directory
	err = bls.checkDotEnv(rootPath)
	if err != nil {
		return err
	}
	//read .env files
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}
	//create loggers and atacched to blacksmith type struct
	infoLog, errorLog := bls.startLoggers()

	//connects to database and
	if os.Getenv("DATABASE_TYPE") != "" {
		//db will be the connection pool
		db, err := bls.OpenDB(os.Getenv("DATABASE_TYPE"), bls.BuildDSN())
		if err != nil {
			errorLog.Println(err)
			os.Exit(1)
		}
		//databse Type and connections pool for myApp instance
		bls.DB = Database{
			DataType: os.Getenv("DATABASE_TYPE"),
			Pool:     db,
		}
	}

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
		//cookies config from .env to manage session
		cookie: cookieConfig{
			name:     os.Getenv("COOKIE_NAME"),
			lifetime: os.Getenv("COOKIE_LIFETIME"),
			persist:  os.Getenv("COOKIE_PERSISTS"),
			secure:   os.Getenv("COOKIE_SECURE"),
			domain:   os.Getenv("COOKIE_DOMAIN"),
		},
		sessionType: os.Getenv("SESSION_TYPE"),
		database: databaseConfig{
			database: os.Getenv("DATABASE_TYPE"),
			dsn:      bls.BuildDSN(),
		},
	}

	//set Session struct type with the paramaters configured in bls.config.cookie
	sess := session.Session{
		CookieLifetime: bls.config.cookie.lifetime,
		CookiePersist:  bls.config.cookie.persist,
		CookieName:     bls.config.cookie.name,
		CookieDomain:   bls.config.cookie.domain,
		SessionType:    bls.config.sessionType,
	}
	//create  SessionManager package type configure with the Session type values
	bls.Session = sess.InitSession()

	//config Jet to initialize the blacksmith.JetViews field
	var views = jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
		jet.InDevelopmentMode(),
	)
	//set the initialized Jet type to blacksmith.JetViews field
	bls.JetViews = views

	// set the renderer for myapp (can be Jet or Go templates depends on .env config)
	bls.createRenderer()
	//if calling he OLD-VERSION //bls.Render = bls.createRenderer(bls)
	return nil
}

//Init loops to folderNames property of initPath struct and creates the webApp folder structure
func (bls *Blacksmith) Init(p initPaths) error {
	//rootPath type contains the root and the folders structure for myapp
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

//ListenAndServe config and start the webserver using myApp *Application.Routes
//(not the blacksmith.routes default handler) as handler (*chi mux) to define
// their own routes in myapp/routes, myApp is overwritten the blacksmith default handler
func (bls *Blacksmith) ListenAndServe() {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", bls.config.port),
		ErrorLog:     bls.ErrorLog,
		Handler:      bls.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}
	//close the DB pool when web server is stopped (make stop command)
	defer bls.DB.Pool.Close()

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

//OLD-VERSION:createRenderer config anc creates a Render struct to atach to blacksmith struct
// func (bls *Blacksmith) createRenderer(b *Blacksmith) {
// 	myRenderer := render.Render{
// 		Renderer: b.config.renderer,
// 		RootPath: b.RootPath,
// 		Port:     b.config.port,
// 	}
// 	return &myRenderer
// }

//this is an alternative way to above function, using only the receivers params to create Renderer type
//and attach the renderer type into blacksmith type, must much cleaner
func (b *Blacksmith) createRenderer() {
	myRenderer := render.Render{
		Renderer: b.config.renderer,
		RootPath: b.RootPath,
		Port:     b.config.port,
		JetViews: b.JetViews,
	}
	b.Render = &myRenderer
}

//BuildDSN build the DB string connection using the values from myapp .env file
func (b *Blacksmith) BuildDSN() string {
	var dsn string
	switch os.Getenv("DATABASE_TYPE") {
	case "postgres", "postgresql":
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5",
			os.Getenv("DATABASE_HOST"),
			os.Getenv("DATABASE_PORT"),
			os.Getenv("DATABASE_USER"),
			os.Getenv("DATABASE_NAME"),
			os.Getenv("DATABASE_SSL_MODE"))

		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("%s password=%s", dsn, os.Getenv("DATABASE_PASS"))
		}
	default:
	}
	return dsn
}
