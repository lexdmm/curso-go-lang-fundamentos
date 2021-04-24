package main // define principal pacote da aplicação.

import (
	"fmt"
	"reflect"
)

func main() {
	var nome = "Lex"
	var versao = 1.1
	var idade = 43
	teste := 0.99

	fmt.Println("OLá Senhor", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("O tipo da variável versao é", reflect.TypeOf(versao))
	fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
	fmt.Println("O tipo da variável idade é", reflect.TypeOf(idade))
	fmt.Println("A teste variável de nome curto foi inferida sem o var, quando utiliza := está implicito que é uma declaração de variável, o tipo dela é", reflect.TypeOf(teste))
}
