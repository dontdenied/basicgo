package main
import ("fmt"
)

func main(){
	/*var nome string
	var idade int
	var cidade string*/
	nomes := []string{"julio", "marcos", "cesar"}
	/*fmt.Print("digite seu nome: ")
	fmt.Scan(&nome)

	fmt.Print("digite sua idade: ")
	fmt.Scan(&idade)

	fmt.Print("digite sua cidade: ")
	fmt.Scan(&cidade)
	
	fmt.Println("nome:", nome, "\nidade:", idade, "\ncidade:", cidade)
	*/
	//printar := strings.Join(nomes, "\n")
	for index, nome:= range nomes{
		fmt.Println(index, nome)
		//fmt.Println(printar)
	}
}