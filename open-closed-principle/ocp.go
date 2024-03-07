package ocp

import "fmt"

// aqui podemos entender que o codigo deve ser fechamos para modificacoes e aberto para extensao
// usando nossa interface
type Empresa interface {
	calculoLucro() float64
}

type Banco struct {
	totalDespesas float64
	totalDeLucros float64
}

type SociedadeCredito struct {
	totalDespesas float64
	totalDeLucros float64
	emprestimos   float64
}

func (b *Banco) calculoLucro() float64 {
	return b.totalDespesas - b.totalDeLucros
}

func (s *SociedadeCredito) calculoLucro() float64 {
	return (s.emprestimos + s.totalDespesas) - s.totalDeLucros
}

func ocp() {

	empresas := []Empresa{&Banco{10, 20}, &SociedadeCredito{10, 20, 30}}

	for _, emempresa := range empresas {
		fmt.Print(emempresa.calculoLucro())
	}

}
