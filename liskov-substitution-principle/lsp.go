package lsp

// garamte que as subclasses podem ser u8sadas em lugares onde as classes base
// sem causar mudancas de comportamento inesperadas

import "fmt"

type Conta struct {
	ContaID   int
	ContaName string
	Saldo     float64
}

type Transacao struct {
	TransacaoID int
	Quantidade  float64
}

func (a *Conta) Sacar(quantidade float64) {
	if quantidade > 0 && quantidade <= a.Saldo {
		a.Saldo -= quantidade
		fmt.Printf("Saque feito $%.2f com sucesso. Novo Saldo: $%.2f\n", quantidade, a.Saldo)
	} else {
		fmt.Println("Saque sem sucesso. Saldo insuficiente.")
	}
}

type ChecaConta struct {
	Conta
}

// Overrides o metodo Saque em tempo de execucao
func (c *ChecaConta) Sacar(quantidade float64) {
	if quantidade > 0 && quantidade <= c.Saldo-100 {
		// Checking se ac conta tem pelo menos 100 de saldo
		c.Saldo -= quantidade
		fmt.Printf("Saque feito $%.2f com sucesso. Novo Saldo: $%.2f\n", quantidade, c.Saldo)
	} else {
		fmt.Println("Saque sem sucesso. Saldo insuficiente.")
	}
}

// metodo deposita conta
type DepositaConta struct {
	Conta
	InterestRate float64
}

func (s *DepositaConta) Sacar(quantidade float64) {
	if quantidade > 0 && quantidade <= s.Saldo {
		// executa a operacao caso correponda as codicionais.
		s.Saldo -= quantidade
		fmt.Printf("Saque feito $%.2f com sucesso. Novo Saldo: $%.2f\n", quantidade, s.Saldo)
	} else {
		fmt.Println("Saque sem sucesso. Saldo insuficiente.")
	}
}

func lsp() {

	checaConta := ChecaConta{
		Conta: Conta{
			ContaID:   1,
			ContaName: "Thyago Costa",
			Saldo:     500,
		},
	}

	depositaConta := DepositaConta{
		Conta: Conta{
			ContaID:   2,
			ContaName: "Mateus Rocha",
			Saldo:     1000,
		},
		InterestRate: 0.02,
	}

	// overide no metodo especializado sacar
	checaConta.Sacar(200)

	// overide no metodo especializado sacar
	depositaConta.Sacar(300)
}
