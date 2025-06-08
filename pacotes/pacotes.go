package pacotes

import (
	"fmt"
	steph "strings"
)

func MostrarPacotes() {
	fmt.Println("Hello, world!")

	// sem alterar o nome do pacote,
	// usaríamos chamando strings.Func()
	// strings.Slipt()
	fmt.Println(steph.Split("steph", ""))
}
