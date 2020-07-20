package main

import (
	"fmt"
)

type humanos interface {
	falar()
}

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) falar() {
	fmt.Printf("Meu é %v e tenho %d anos.", p.nome, p.idade)
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {

	p := pessoa{" Allan santos", 80}

	
	dizerAlgumaCoisa(&p)
}
