package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
 Criar uma (1) função e seu protótipo: PESQUISAR, que receba como argumento uma mensagem e
uma palavra e retorne a Quantidade de Vezes (QV) que a palavra ocorre na mensagem.
No programa principal (main), use as funções criadas em 100 (mensagens + palavras) lidas do
usuário (uma por uma). Considere letras do alfabeto: 'A' até 'Z' (Maiúsculas) e 'a' até 'z' (Minúsculas)
*/

func main() {

	var palavra string

	var letra byte

	var entradaPalavra string

	contador := 0

	fmt.Println("Insira a frase")

	in := bufio.NewReader(os.Stdin)
	palavra, err := in.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Insira a letra que deseja pesquisar")

	fmt.Scan(&entradaPalavra)

	letra = entradaPalavra[0]

	for i := 0; i < len(palavra); i++ {

		if (palavra[i]) == letra {
			contador++
		}

	}

	fmt.Println("A letra aparece", contador, "vezes")
}
