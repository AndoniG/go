package main

import (
	"fmt" 
	"net/http"
	"io"
)

type webWritter struct {

}

func (webWritter) Write (p []byte) (int, error){
	fmt.Println(string(p))
	return len(p), nil
}

func main(){
	response, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}
	e := webWritter{}
	// COPY REQUIERE UNA INTERFAZ DE ESCRITOR QUE ES UN STRUC CON LA FUNCIÓN WRITE
	io.Copy(e, response.Body)
}