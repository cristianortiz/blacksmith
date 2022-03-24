package session

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
)

// Session type for SessionManager package to handle session in myapp also allows to write
// an read sessions from DB with sessions store por every single DB engine
type Session struct {
	CookieLifetime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	CookieSecure   string
	SessionType    string
	DBPool         *sql.DB //to write to the session table and read from it
}

//InitSession set up sessionManager type form the .env file, as every value in it is a string
//must be converted into the types requierd by sessionManager package
func (c *Session) InitSession() *scs.SessionManager {
	var persist, secure bool

	//how long should session last
	minutes, err := strconv.Atoi(c.CookieLifetime)
	if err != nil {
		minutes = 60
	}
	//shold cookies persist
	if strings.ToLower(c.CookiePersist) == "true" {
		persist = true
	}
	//should cookies secure
	if strings.ToLower(c.CookieSecure) == "true" {
		secure = true
	}
	//create session
	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Name = c.CookieDomain
	session.Cookie.Secure = secure
	session.Cookie.Domain = c.CookieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode

	// config the DBPool for write and read session
	// from DB postgres, mysql,redis etc using sessionManager session store
	switch strings.ToLower(c.SessionType) {

	case "redis":
	case "mysql", "mariadb":
		session.Store = mysqlstore.New(c.DBPool)
	case "postgres", "postgresql":

		session.Store = postgresstore.New(c.DBPool)
	default:
		//cookie

	}
	return session
}
