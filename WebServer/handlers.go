package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func HandleRoot(w http.ResponseWriter, r *http.Request){
	// FPRINT RECIBE UN WRITER
	fmt.Fprintf(w, "Hello World")
}

func HandleHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "API endpoint")
}

func PostRequest(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var metadata MetaData
	err := decoder.Decode(&metadata)
	if err != nil {
		// %v permite enviar texto legible
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "Payload %v\n", metadata)
}

func UserPostRequest(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		// %v permite enviar texto legible
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	obj, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "applications/json")
	w.Write(obj)
}