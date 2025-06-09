package maps

import "fmt"

func MostrarMaps() {
	idade := map[string]int{}
	idade["tata"] = 28
	idade["bento"] = 4

	fmt.Printf("%v\n", idade)
	fmt.Printf("%v\n", idade["tata"])
	fmt.Printf("%v\n", idade["bento"])
}
