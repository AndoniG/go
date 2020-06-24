package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	start := time.Now()
	// CREAMOS UN CANAL QUE ESTARÁ TRANSMITIENDO STRINGS
	channel := make(chan string)
	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}


	// for _, server := range servers{
	// 	go checkServer(server, channel)
	// }
	
	// REQUERIMOS ESPECIFICAR EXPLÍCITAMENTE LOS PROCESOS A ESPERAR
	// for i:=0 ; i < len(servers); i++ {
	// 	//  <- channel envía información
	// 	fmt.Println(<- channel)
	// }

	i := 0
	for {
		if i > 10{
			break
		}
		for _, server := range servers{
			go checkServer(server, channel)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<- channel)
		i++
	}

	done := time.Since(start)
	fmt.Printf("Tiempo de ejecución %s\n", done)
}

func checkServer(server string, channel chan string){
	_, err := http.Get(server)
	if err != nil{
		// chan <- INDICA QUE EL CANAL RECIBE INFORMACIÓN
		channel <- server + " no se encuentra disponible"
	} else {
		channel <- server + " está funcionando correctamente"
	}
}