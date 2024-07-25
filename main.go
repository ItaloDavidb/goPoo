package main

import "fmt"

type Conta struct {
	nome  string
	saldo float64
}

func main() {
	conta1 := Conta{nome: "teste", saldo: 35.2} // var de tipo valor
	conta2 := Conta{nome: "teste2", saldo: 30}  // var de tipo valor
	conta3 := new(Conta)                        // criando dessa forma acaba criando um ponteiro
	conta3.nome = "teste3"
	conta3.saldo = 44
	fmt.Println(conta1)
	conta1.deposito(30)
	fmt.Println(conta1)
	fmt.Println("Teste Deposito")
	conta1.transferencia(65.2, &conta2)
	fmt.Println(conta1)
	fmt.Println(conta2)
	fmt.Println(&conta3)
	fmt.Println(*conta3)
	fmt.Println("Teste Transferencia")
}

func (c *Conta) deposito(value float64) {
	c.saldo += value
}
func (c *Conta) transferencia(value float64, contaTransferir *Conta) {
	c.saldo -= value
	contaTransferir.deposito(value)
}

// & == endereco
// * == conteudo dentro do endereco
