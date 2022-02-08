package session

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
)

type Session struct {
	CookieLifetime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	CookieSecure   string
	SessionType    string
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

	//wich session store
	switch strings.ToLower(c.SessionType) {

	case "redis":
	case "mysql", "mariadb":
	case "postgres", "postgresql":
	default:
		//cookie

	}
	return session
}
