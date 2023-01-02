package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Ejecucion finalizada")
	}()

	file, err := os.Open("customers.txt")
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic("No se puede leer el archivo")
	}
	fmt.Println(string(data))
	file.Close()
}
