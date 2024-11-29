package main
import "fmt"

func main(){
	//cbf := make([]int, 0, 3)
	//lista em go
	ferj := [] int{0, 1, 2, 3, 4, 5}
	//iteração em go
	for i:= range ferj{
		fmt.Println(i)
	}
	fmt.Println(ferj)
	// dicionário em go
	tabela:= map[string]int{
		"flamengo":91,
		"santos":82,
		"palmeiras": 72,

	}
	// iteração de dicionario em go
	for chave, valor:= range tabela{
		//cbf = append(chave)
		
		
		fmt.Println("time:", chave, "pontos", valor)
	}
	//fmt.Println(cbf)



}