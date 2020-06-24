package main

import (
	"fmt"
	"net/http"
	"time"
	"log"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request){
			flag := true
			fmt.Println("Authenticating User")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request){
			start := time.Now()
			defer func(){
				// LOG ACCEDE AL SISTEMA
				log.Println(r.URL.Path, time.Since(start))
			}() //EJECUCIÓN DE FUNCIÓN ANÓNIMA
			f(w, r)
		}
	}
}