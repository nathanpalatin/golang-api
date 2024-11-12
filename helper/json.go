package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequest(req *http.Request, result interface{})  {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(result)
	ErrorPanic(err)
}

func WriteResponse(w http.ResponseWriter, data interface{}) {
		w.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		err := encoder.Encode(data)
		ErrorPanic(err)
}