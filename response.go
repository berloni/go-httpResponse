package response

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// Response defines the base successfull response of an handler
type Response struct {
	Err        bool        `json:"err"`
	ErrMsg     string      `json:"errMsg"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

// Error creates an error response for an handler
func Error(w *http.ResponseWriter, err error, statusCode int) {
	var response Response
	(*w).Header().Set("Content-Type", "application/json")

	response.Err = true
	response.ErrMsg = err.Error()
	response.StatusCode = statusCode
	response.Data = nil

	payload, _ := json.Marshal(response)
	(*w).Write([]byte(payload))
}

// Success creates n success response for an handler
func Success(w *http.ResponseWriter, data interface{}, statusCode int) {
	var response Response
	(*w).Header().Set("Content-Type", "application/json")

	response.Err = false
	response.ErrMsg = ""
	response.StatusCode = statusCode
	response.Data = data

	payload, _ := json.Marshal(response)
	(*w).Write([]byte(payload))
}

// HTMLPage shows an html page as response for an handler
func HTMLPage(w *http.ResponseWriter, t *template.Template, tVars interface{}) {
	(*w).Header().Set("Content-Type", "text/html")
	htmlErrPage := "<h2>Unknown error, please try again</h2>"

	if t != nil {

		err := t.Execute((*w), tVars)
		if err != nil {
			fmt.Fprint((*w), htmlErrPage)
			return
		}

	} else {
		fmt.Fprint((*w), htmlErrPage)
	}
}
