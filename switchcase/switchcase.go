package switchcase

import (
	"fmt"
)

func MostarSwitchcase() {

	posicao := 2
	switch posicao {
	case 1:
		fmt.Println("Primeiro lugar")
	case 2:
		fmt.Println("Segundo lugar")
	case 3:
		fmt.Println("Terceiro lugar")
	}
}
