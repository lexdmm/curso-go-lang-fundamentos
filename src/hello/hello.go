package main

import "fmt"

func main() {
	nome := "Lex"
	idade := 24
	versao := 1.1
	fmt.Println("Olá sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scan(&comando) // &comercial define o ponteiro que vai receber o dado desta variável, não uso scanf e sim scan porque eu ja sei que é inteiro.
	fmt.Println("O comando escolhido foi", comando, "Endereço da variável", &comando)
}
