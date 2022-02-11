package session

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alexedwards/scs/v2"
)

func TestSession_InitSession(t *testing.T) {
	//mock session data for a cookie
	c := &Session{
		CookieLifetime: "100",
		CookiePersist:  "true",
		CookieName:     "blacksmith",
		CookieDomain:   "localhost",
		SessionType:    "cookie",
	}
	//SessionManager type variable pointer
	var sm *scs.SessionManager

	//this var should be a SessionManager type
	ses := c.InitSession()

	var sessKind reflect.Kind
	var sessType reflect.Type

	rv := reflect.ValueOf(ses)

	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		fmt.Println("for loop:", rv.Kind(), rv.Type(), rv)
		sessKind = rv.Kind()
		sessType = rv.Type()

		rv = rv.Elem()
	}
	if !rv.IsValid() {
		t.Error("invalid type or kind; kind", rv.Kind(), "type:", rv.Type())
	}
	if sessKind != reflect.ValueOf(sm).Kind() {
		t.Error("wrong kind returned testing cookie session, Expected", reflect.ValueOf(sm).Kind(), " and got ", sessKind)
	}

	if sessType != reflect.ValueOf(sm).Type() {
		t.Error("wrong type returned testing cookie session, Expected", reflect.ValueOf(sm).Type(), " and got ", sessType)
	}

}
