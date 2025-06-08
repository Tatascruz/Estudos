package tipos

import (
	"fmt"
)

func MostrarTipos() {
	// bool (true/false)
	fmt.Printf("Type: %T - Value %v\n", true, true)

	//string - sequência de bytes
	fmt.Printf("Type: %T - Value %v\n", "Tata", "Tata")
	fmt.Printf("Type: %T - Value %v\n", "1", "1")

	// int
	fmt.Printf("Type: %T - Value %v\n", 1, 1)

	// float (float64/float32) - decimal
	fmt.Printf("Type: %T - Value %v\n", 1.233, 1.233)
}
