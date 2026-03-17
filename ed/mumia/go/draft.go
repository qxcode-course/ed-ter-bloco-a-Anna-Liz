package main

import "fmt"

func main() {
	var nome string
	var idade int
	var classificacao string

	fmt.Scan(&nome, &idade)

	switch {
	case idade < 12:
		classificacao = "crianca"
	case idade >= 12 && idade < 18:
		classificacao = "jovem"
	case idade >= 18 && idade < 65:
		classificacao = "adulto"
	case idade >= 65 && idade < 1000:
		classificacao = "idoso"
	default:
		classificacao = "mumia"
	}

	fmt.Println(nome + " eh " + classificacao)

}
