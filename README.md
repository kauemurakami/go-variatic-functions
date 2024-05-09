[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-variatic-functions/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-variatic-functions/blob/main/README.md)  
go version 1.22.1

## Variatic Functions
It is basically a function that can receive ```n``` parameters, without needing to specify the parameters.  

### Starting project
Create a ```go-variatic-functions``` directory with a ```main.go``` file, with the following structure:  
```go
package main

func main() {

}
```

### Using variatic functions
As an example, let's create a function that can add ```n``` numbers, that is, we will no longer specify that it will receive 2 numbers, for example, I just want to pass 4 or more, or less:  
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
Here we can see that ```numbers``` is a ```slice``` of type ```int```, which can receive as many values ​​as you want to insert, as the size is not defined.  
To sum the past numbers, we use an auxiliary variable ```total``` that receives 0 initially.
We use ```for``` to retrieve a ```number``` from the ```numbers``` slice, ignoring the ```index``` with ```_``` at the beginning, for each number in the ```slice```, it adds up to the value contained in ```total```.<br/><br/>

*Variatic function with fixed parameters*  
```go
...

func write(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
...
	write("See the number", 1, 2, 3, 4, 5, 6, 7)
	// output
	//See the number 1
	//See the number 2
	//See the number 3
	//See the number 4
	//See the number 5
	//See the number 6
	//See the number 7

}
```
Here we have the same behavior as the ```sum``` function, but this time we add a unique ```text``` variable, and a ```slice``` of ```int``` again, and we will go through it with ```range```, always displaying the message and each ```number``` found in the ```numbers``` slice.
