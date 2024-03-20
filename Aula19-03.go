package main

import (
	"fmt"
)

func main() {

	var opcao int
	fmt.Println(`
	████████████████████████████████████████████████████████████████████████████
	█░░░░░░░░░░░░░░█░░░░░░░░░░░░░░█░░░░░░░░░░░░░░█░░░░░░░░░░░░░░█░░░░░░░░░░░░░░█
	█░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█
	█░░▄▀░░░░░░░░░░█░░▄▀░░░░░░▄▀░░█░░░░░░▄▀░░░░░░█░░▄▀░░░░░░░░░░█░░▄▀░░░░░░░░░░█
	█░░▄▀░░█████████░░▄▀░░██░░▄▀░░█████░░▄▀░░█████░░▄▀░░█████████░░▄▀░░█████████
	█░░▄▀░░░░░░░░░░█░░▄▀░░░░░░▄▀░░█████░░▄▀░░█████░░▄▀░░░░░░░░░░█░░▄▀░░█████████
	█░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█████░░▄▀░░█████░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀░░█████████
	█░░▄▀░░░░░░░░░░█░░▄▀░░░░░░▄▀░░█████░░▄▀░░█████░░▄▀░░░░░░░░░░█░░▄▀░░█████████
	█░░▄▀░░█████████░░▄▀░░██░░▄▀░░█████░░▄▀░░█████░░▄▀░░█████████░░▄▀░░█████████
	█░░▄▀░░█████████░░▄▀░░██░░▄▀░░█████░░▄▀░░█████░░▄▀░░░░░░░░░░█░░▄▀░░░░░░░░░░█
	█░░▄▀░░█████████░░▄▀░░██░░▄▀░░█████░░▄▀░░█████░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█
	█░░░░░░█████████░░░░░░██░░░░░░█████░░░░░░█████░░░░░░░░░░░░░░█░░░░░░░░░░░░░░█
	████████████████████████████████████████████████████████████████████████████`)

	for opcao != 99 {
		fmt.Println("MENU ATIVIDADE")
		fmt.Println("1. atividade 1")
		fmt.Println("2. atividade 2")
		fmt.Println("3. Atividade 9")
		fmt.Println("4. Atividade 10")
		fmt.Println("5. Atividade 11")
		fmt.Println("6. Atividade 12")
		fmt.Println("99. Finalizar/Encerrar")

		fmt.Print("Insira a opção desejada: ")
		fmt.Scanln(&opcao)

		if opcao == 1 {
			// Exercicio 1
			arr1 := [5]int{1, 2, 3, 4, 5}
			fmt.Println(arr1)
			opcao = 99
		}

		if opcao == 2 {
			arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
			fmt.Println(arr2)
			opcao = 99
		}

		if opcao == 3 {
			// Exercicio 9

			fmt.Println("")

			arr9 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

			fmt.Println(arr9)

			for i := 0; i < 10; i++ {

				if arr9[i]%2 == 0 {
					fmt.Println(arr9[i], "é um número par")
				}
			}
		}
		if opcao == 4 {
			// Exercicio 10

			fmt.Println("")

			arr10 := [5]int{1, 3, 5}

			fmt.Println(arr10)

			fmt.Println("Insira os valores que faltam!")
			for i := 0; i < 5; i++ {
				if arr10[i] == 0 {
					var item int
					fmt.Printf("Valor / Posição %d: ", i)
					fmt.Scanln(&item)
					arr10[i] = item

				}
			}
			fmt.Println(arr10)
		}

		if opcao == 5 {
			// Exercicio 11

			fmt.Println("")

			arr11 := [5]int{}
			arr11Copia := [5]int{}
			var valorAdicionar int
			var valorExcluir int

			fmt.Println("Insira os valores que deseja em seu array!")

			for z := 0; z < 5; z++ {

				fmt.Printf("Valor %d: ", z)
				fmt.Scanln(&valorAdicionar)
				arr11[z] = valorAdicionar
				arr11Copia[z] = valorAdicionar
			}

			fmt.Println("Informe o valor que deseja excluir!")
			fmt.Scanln(&valorExcluir)

			for c := 0; c < 5; c++ {
				if arr11[c] == valorExcluir {
					arr11[c] = 0

				}
			}

			fmt.Println(arr11, "Com o valor excluido")
			fmt.Println(arr11Copia, "Sem o valor excluido")
		}

		if opcao == 6 {
			// exercicio 12

			fmt.Println("")

			valorProcurar := 0
			var valorEncontrado int
			arr12 := [5]int{1, 2, 3, 4, 5}
			fmt.Println("Insira o valor que deseja procurar no array a seguir: ", arr12)
			fmt.Scanln(&valorProcurar)

			for v := 0; v < 5; v++ {
				if arr12[v] == valorProcurar {
					fmt.Printf("Valor %d, encontrado no indice %d", valorProcurar, v)
					valorEncontrado = 1
				}
			}
			if valorEncontrado == 0 {
				fmt.Printf("Valor não encontrado.")
			}
		}
	}
}
