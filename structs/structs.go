package structs

import "fmt"

// Structs
// Forma de criar sua própria estrutura de dados
// Personalizar de acordo com a sua necessidade
// Podemos usar vários tipos diferentes

// type <nome da nossa estrutura> struct { <campos> }
type Pessoa struct {
	Nome  string
	Idade int
}

type Profissao struct {
	Pessoa
	Tipo string
}

func MostrarStructs() {
	  p := Pessoa{Nome: "tata", Idade: 39}
	  fmt.Println(p)
	  fmt.Println(Pessoa{Nome: "tata"})

	p1 := Pessoa{Nome: "Bob", Idade: 2}
	  fmt.Println(p1.Nome)
	  fmt.Println(p1.Idade)

	  p1.Idade = 3
	  fmt.Println(p1.Idade)

	p2 := Pessoa{Nome: "Patrick", Idade: 2}

	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	  fmt.Println(pessoas)

	// structs + map
	  alunos := map[string][]Pessoa{}
	  ["programação"] = pessoas
	fmt.Println(alunos)

	var alunos = map[string][]Pessoa{
	  "programação": {{Nome: "tata", Idade: 39}, {Nome: "bento", Idade: 4}},
	  "engenharia":  {{Nome: "tata2", Idade: 39}, {Nome: "bento2", Idade: 4}},
	}
	fmt.Println(alunos)

	struct herdando  campos de outro struct
	prof := Profissao{p2, "dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Nome)
	fmt.Println(prof.Pessoa.Idade)
	fmt.Println(prof.Tipo)
}
