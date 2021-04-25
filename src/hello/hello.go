package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome, _ := devolveNomeEIdade() //Posso utilizar somente um dos retornos
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando) // &comercial define o ponteiro que vai receber o dado desta variável, não uso scanf e sim scan porque eu ja sei que é inteiro.
	fmt.Println("O comando escolhido foi", comando, "Endereço da variável", &comando)

	return comando
}

func devolveNomeEIdade() (string, int) {
	nome := "Lex"
	idade := 43
	return nome, idade
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://www.google.com"
	resp, _ := http.Get(site)
	fmt.Println(resp)
}
