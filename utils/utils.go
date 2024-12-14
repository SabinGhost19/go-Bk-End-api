package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

var Validator=validator.New();

func ParseJson(req *http.Request, payload any)error{
	if req.Body==nil{
		return fmt.Errorf("missing request body");
	}
	return json.NewDecoder(req.Body).Decode(&payload);
}

func WriteJson(w http.ResponseWriter,status int,what_to_write any)error{

	w.Header().Add("Content-Type","application/json");
	w.WriteHeader(status);
	return json.NewEncoder(w).Encode(what_to_write);
} 
func WriteJsonError(w http.ResponseWriter,status int,err error)error{
	return WriteJson(w,status,map[string]string{"error":err.Error()});
}

