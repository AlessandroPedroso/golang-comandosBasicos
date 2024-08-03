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

	// nome, idade := devolveNomeEIdade()
	// fmt.Println(nome, idade)

	// retornar a variavel que você aquiser
	// _, idade := devolveNomeEIdade()
	// fmt.Println(idade)

	switch comando {
	case 1:

		iniciarMonitoramento()

	case 2:
		fmt.Println("Exibindo logs")

	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)

	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}

}

// exemplo de retorno com duas funções
func devolveNomeEIdade() (string, int) {
	nome := "Douglas"
	idade := 29

	return nome, idade
}

func exibeIntroducao() {
	nome := "Alessandro"
	versao := 1.1

	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("\n1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0 - Saír do Programa")
}

func leComando() int {

	var comandoLido1 int
	fmt.Scan(&comandoLido1)

	fmt.Println("O comando escolhido foi", comandoLido1)

	return comandoLido1
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando")
	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)
	fmt.Println(resp)

}

// var nome string = "Alessandro"
// var idade int = 29
// var versao float32 = 1.1

//sem declaração de tipos de variaveis
// var nome = "Alessandro"
// var idade = 29

// float precisa declarar pois tem duas opções de tipos
// var versao float32 = 1.1

//declarando variaveis sem informar o VAR

// idade := 29

// if comando == 1 {
// 	fmt.Println("Monitorando")
// } else if comando == 2 {
// 	fmt.Println("Exibindo logs")
// } else if comando == 0 {
// 	fmt.Println("Saindo do programa")
// } else {
// 	fmt.Println("Não conheço este comando")
// }

// fmt.Scanf("%d", &comando)

//reflect para saber o tipo da variavel
// fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
