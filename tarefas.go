package main
import "fmt"


func main(){
	adicionar_tarefa()

}

func adicionar_tarefa(){
	diciona := make(map[string]interface{})
	var nome string
	var descricao string
	var prioridade int
	

	fmt.Print("digite o nome da tarefa: ")
	fmt.Scanln(&nome)
	
	fmt.Print("digite a descrição da tarefa: ")
	fmt.Scanf("%q", &descricao)

	fmt.Print("digite a prioridade 1/5: ")
	fmt.Scanln(&prioridade)



	diciona["nome"] = nome
	diciona["descriçao"] = descricao
	diciona["prioridade"] = prioridade

	

	for chave, valor:= range diciona{
		fmt.Println(chave, valor)
	}

}


