[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-variatic-functions/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-variatic-functions/blob/main/README.md)  
go version 1.22.1

## Funções Variáticas
É basicamente uma função que pode receber ```n``` parâmentros, não necessitando de especificar os parametros.  

### Iniciando projeto
Crie um diretório ```go-variatic-functions``` com um arquivo ```main.go```, com a seguinte estrutura:  
```go
package main

func main() {

}
```

### Usando as funções variáticas
Como exemplo vamos criar uma função que pode somar ```n``` números, ou seja, não vamos mais especificar que ela vai receber 2 números por exemplo apenas, quero passar 4 ou mais, ou menos:  
```go
....
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func write(text string, numbers ...int){
  for _, number := range numbers {
    fmt.Println(text, number)
  }
}

func main() {
	fmt.Println(sum(2, 5, 3))           //output 10
	fmt.Println(sum(2, 5, 3, 5, 8, 10)) //output 33
  
}
```
Aqui podemos ver que ```numbers``` é um ```slice``` do tipo ```int```, que pode receber quantos valores você quiser inserir, por não ter tamanha definido.  
Para a soma dos números passados usamos uma variável auxiliar ```total``` que recebe 0 inicialmente.  
Usamos o ```for``` para recuperar um ```number``` do slice de ```numbers```, ignorando o ```index``` com ```_``` no início, a cada número no ```slice```, ele soma com o valor contido em ```total```.<br/><br/>

*Função variáticas com parâmetros fixos*  
```go
...

func write(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
...
	write("Olha o número", 1, 2, 3, 4, 5, 6, 7)
	// output
	//Olha o número 1
	//Olha o número 2
	//Olha o número 3
	//Olha o número 4
	//Olha o número 5
	//Olha o número 6
	//Olha o número 7

}
```
Aqui temos o mesmo comportamento da função ```sum```, mas dessa vez adicionamos uma variavel ```text```, única, e um ```slice``` de ```int``` novamente, e vamos percorre-lo com ```range```, exibindo sempre a mensagem e cada ```number``` encontrado no slice de ```numbers```.
