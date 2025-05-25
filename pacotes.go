package main

import (
	"fmt"
	steph "strings"
)

func mostrarPacotes() {
	fmt.Println("Hello, world!")

	// sem alterar o nome do pacote,
	// usaríamos chamando strings.Func()
	// strings.Slipt()
	fmt.Println(steph.Split("steph", ""))
}
