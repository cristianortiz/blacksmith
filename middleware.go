package blacksmith

import "net/http"

func (bls *Blacksmith) SessionLoad(next http.Handler) http.Handler {
	//check if the sessionLoad midlware is called on myapp init
	bls.InfoLog.Println("==> SessionLoad called")
	return bls.Session.LoadAndSave(next)
}
