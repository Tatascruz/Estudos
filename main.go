package main

import (
	"estudos/funcoes"
	ifelse "estudos/if-else"
	"estudos/listas"
	"estudos/loops"
	"estudos/maps"
	"estudos/pacotes"
	"estudos/structs"
	"estudos/switchcase"
	"estudos/tipos"
)

func main() {
	listas.MostrarArray()
	tipos.MostrarTipos()
	funcoes.Saudacao("Tata")
	pacotes.MostrarPacotes()
	maps.MostrarMaps()
	structs.MostrarStructs()
	ifelse.MostraIfElse()
	switchcase.MostarSwitchcase()
	loops.MostrarLoops()

}
