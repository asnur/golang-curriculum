package main

import "fmt"

func main() {
	var number int = 5               //deklarasi variabel number dengan nilai 5
	var numberPointer *int = &number //deklarasi variabel numberPointer sebagai pointer dari variabel number

	fmt.Println("Nilai dari number:", number)
	fmt.Println("Alamat memory dari number:", &number)
	fmt.Println("Nilai dari numberPointer:", *numberPointer)
	fmt.Println("Alamat memory dari numberPointer:", numberPointer)
}
