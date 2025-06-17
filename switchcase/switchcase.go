package switchcase

import (
	"fmt"
)

func MostarSwitchcase() {

	// posicao := 2
	// switch posicao {
	// case 1:
	//	fmt.Println("Primeiro lugar")
	// case 2:
	//	fmt.Println("Segundo lugar")
	// case 3:
	//	fmt.Println("Terceiro lugar")
	// }

	nome := "bob"
	switch nome {
	case "tata":
		fmt.Println("É uma pessoa")
	case "bento":
		fmt.Println("É um cachorro")
	case "bob":
		fmt.Println("É um personagem")
	}
}
