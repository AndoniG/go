package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	start := time.Now()
	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, server := range servers{
		checkServer(server)
	}

	done := time.Since(start)
	fmt.Printf("Tiempo de ejecución %s\n", done)
}

func checkServer(server string){
	_, err := http.Get(server)
	if err != nil{
		fmt.Println(server, "no está disponible =(")
	} else {
		fmt.Println(server, "está funcionando correctamente")
	}
}