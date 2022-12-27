package main

import "fmt"

func main() {
	var palabra string
	fmt.Println("Ingrese una palabra")
	fmt.Scanf("%s", &palabra)
	fmt.Println("Cantidad de letras:", len(palabra))

	//for i := 0; i < len(palabra); i++ {
	//	fmt.Println(string(palabra[i]))
	//}

	for _, letra := range palabra {
		fmt.Println(string(letra))
	}

}
