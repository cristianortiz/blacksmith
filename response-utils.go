package blacksmith

import (
	"encoding/json"
	"net/http"
)

//WriteJSON() serialize JSON utility for return data in htto request, receive a response writer,
//a http status code, any kind of data can be received so is generalize it as and interface,
//and optionally some http headers as variadic parameters
func (bls *Blacksmith) WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	//1- out, err := json.MarshalIndent(data, "", "\t")
	// if err != nil {
	// 	return err
	// }

	//TEST: use encoder instead of json marshall, let see if works
	e := json.NewEncoder(w)

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	e.Encode(data)
	//1- _,err =w.Write(out)
	// if err != nil {
	// 	return err
	// }
	return nil

}
