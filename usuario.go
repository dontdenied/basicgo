package main

import (
	"fmt"
)

func main() {
	// Criando o mapa para armazenar as informações do usuário
	usuario := make(map[string]interface{})

	// Variáveis para armazenar o nome, idade e altura
	var nome string
	var idade int
	var altura float64

	// Solicitar o nome
	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&nome)

	// Solicitar a idade
	fmt.Print("Digite sua idade: ")
	fmt.Scanln(&idade)

	// Solicitar a altura
	fmt.Print("Digite sua altura (em metros): ")
	fmt.Scanln(&altura)

	// Armazenando os dados no mapa
	usuario["nome"] = nome
	usuario["idade"] = idade
	usuario["altura"] = altura

	// Exibindo as informações armazenadas no mapa
	
	for chave, valor:= range usuario{
		fmt.Println("chave:", chave,"valor:", valor)
	}
}
