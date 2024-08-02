package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando")

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

	var comandoLido int
	fmt.Scan(&comandoLido)

	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
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
