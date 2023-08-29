// to unmarshal the json sent by user

package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	// defer r.Body.Close() //
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

// This code defines a function named ParseBody which is designed to read and parse the JSON body from an HTTP request. Let's break down each part of the code:

// func ParseBody(r *http.Request, x interface{}) {

// ParseBody is a function that takes two parameters:
// r *http.Request: A pointer to an HTTP request object that is used to access the request's body.
// x interface{}: An empty interface that can hold values of any type. This is the target where the parsed JSON data will be stored.
// defer r.Body.Close()

// defer is used to schedule the closing of the request body (r.Body) after the function returns. This ensures that the body is properly closed, preventing resource leaks.
// if body, err := ioutil.ReadAll(r.Body); err == nil {

// ioutil.ReadAll(r.Body): Reads the entire request body (r.Body) and returns its content as a byte slice ([]byte).
// body, err := ioutil.ReadAll(r.Body): The result of ioutil.ReadAll is assigned to the body variable, and any error that occurs during reading is assigned to the err variable.
// err == nil: Checks if there was no error while reading the request body.
// if err := json.Unmarshal([]byte(body), x); err != nil {

// json.Unmarshal([]byte(body), x): Attempts to parse (unmarshal) the JSON data from the body byte slice into the value pointed to by x.
// err != nil: Checks if there was an error during the unmarshaling process.
// return

// If an error occurred either during reading the request body or unmarshaling the JSON data, the function immediately returns without performing any further actions.
// Overall, this code is intended to read the JSON body of an HTTP request, unmarshal it into the provided target variable x, and close the request body afterward. However, it lacks proper error handling and doesn't provide any indication of what went wrong if an error occurs. It would be advisable to improve the error handling by returning error values or logging relevant information to help diagnose any issues that might arise during the parsing process.
