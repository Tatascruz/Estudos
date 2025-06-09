package maps

import "fmt"

func MostrarMaps() {
	idade := map[string]int{}
	idade["tata"] = 28
	idade["bento"] = 4

	fmt.Printf("%v\n", idade)
	fmt.Printf("%v\n", idade["tata"])
	fmt.Printf("%v\n", idade["bento"])

	anoNasc := map[string]int{
		"tata":  1985,
		"bento": 2008,
	}
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["tata"])
	fmt.Println(anoNasc["bento"])
	anoNasc["golangDoZero"] = 2024
	fmt.Println(anoNasc)

}
