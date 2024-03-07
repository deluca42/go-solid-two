package srp

import "fmt"

// struct para representar uma pessoa
type Instituicao struct {
	nome string
	ispb int
	tipo string
}

// metodo responsavel por exibir o nome da instituicao
func (i Instituicao) ExibirNome() {
	fmt.Println("Nome instituicao", i.nome)
}

// metodo responsavel por exibir o tipo da instituicao
func (i Instituicao) ExibirTipo() {
	fmt.Println("Tipo instituicao", i.tipo)
}

func (i Instituicao) ExibirIspb() {
	fmt.Println("Ispb instituicao", i.ispb)
}

// cada metodo tem uma responsabilidade, garantindo
// que se o comportamento de exibir nome for alterado
// apenas a alteracao ficara no metodo exibirNome
func srp() {

	instituicao := Instituicao{
		nome: "Banco do Brasil",
		tipo: "Banco varejo",
		ispb: 00000000,
	}

	instituicao.ExibirNome()
	instituicao.ExibirTipo()
	instituicao.ExibirTipo()

}
