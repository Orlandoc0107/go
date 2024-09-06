package main

import (
	"fmt"

	"rsc.io/quote"
)

// main imprime por consola los mensajes de bienvenida y las citas de los
// paquetes importados.
func main() {
	fmt.Println("Hola Mundo")

	fmt.Println(quote.Hello()) // 
	fmt.Println(quote.Glass()) // 
	fmt.Println(quote.Go()) // 
	fmt.Println(quote.Opt()) // 
	
	
	// go obliga que nosotros utilicemos las variables o dara error
	var firstName, lastName string
	var age int 

	firstName ="Orlando"
	lastName ="Cardenas"
	age = 100

	fmt.Println(firstName, lastName, age)


	var ( 
		firstName2 string = "Orlando2"
		lastName2 string = "Cardenas2"
		age2 int = 100
	)

	fmt.Println(firstName2, lastName2, age2)
}

// si queremos copilar el archivo con el comando go build hola.go
// para ejecutar el archivo compilado en linux se ejecuta con el comando go run hola.go
// fmt.PrintIn con este comando se imprimen en la consola sin salto de linea 
// se ejecuta con el comando go run hola.go
// con el comando go mod init hola se crea el modulo de manejador de modulos
// el archivo go.mod es el archivo de configuracion del modulo
// con el comando "go get rsc.io/quote" se descarga el modulo de rsc.io/quote, este es un modulo de apoyo para la consola de go.




