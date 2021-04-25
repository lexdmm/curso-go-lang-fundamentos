package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const quantidadeMonitoramentos = 3
const delay = 5

func main() {
	lerSitesArquivo()
	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
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

	sites := sitesMonitoramento()

	for i := 0; i < quantidadeMonitoramentos; i++ {
		/*
			posso utilizar o range para retornar a posição e o conteúdo do array,
			nesse caso não me interessa a posição então digo que não vou utilizar com _
			mas se fosse utilizar posição seria "for i, site := range sites { ... }"
		*/
		for _, site := range sites {
			testaSitesMonitoramento(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func sitesMonitoramento() []string {
	sites := []string{
		"https://random-status-code.herokuapp.com/",
		"https://www.google.com.br",
		"https://www.whatsapp.com/?lang=pt_br",
		"https://stackoverflow.com/",
	}

	return sites
}

func testaSitesMonitoramento(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro de protocolo GET:", err)
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

func lerSitesArquivo() []string {
	var sites []string
	arquivo, err := ioutil.ReadFile("sites")

	if err != nil {
		fmt.Println("Ocorreu um erro de leitura:", err)
	}
	fmt.Println(string(arquivo))
	return sites
}
