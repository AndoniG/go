package main

import ("fmt")

func main(){
	// CREAMOS UN MAPA DE PROPIEDAD STRING Y VALOR INT
	m1 := make(map[string]int)
	m1["a"] = 8
	fmt.Println(m1)
}