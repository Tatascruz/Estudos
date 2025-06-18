package loops

import (
	"fmt"
)

// Loops
// Laços de repetição
// Repetir tarefas

func MostrarLoops() {

	// sum := 0
	// for i := 1; i < 10; i++ {
	//	fmt.Println("Sum:", sum)
	//	fmt.Println("Indice:", i)
	//	sum += i

	// é a mesma coisa que: sum + i
	// sum -= i -> sum = sum - 1

	// É como se ao final do loop
	// fosse adicionado 1 ao indice i
	// i++
	// i = i + 1
	// }
	// fmt.Println(sum)

	// loop infinito
	// for {
	// fmt.Println("Golang do zero")
	// time.Sleep(2 * time.Second)
	// }

	// // for range
	frutas := []string{"laranla", "maça", "banana", "uva", "kiwi"}
	for _, frutas := range frutas {
		fmt.Println("frutas", frutas)
	}
}
