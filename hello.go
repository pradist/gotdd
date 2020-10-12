package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if len(name) == 0 {
		name = "World"
	}

	prefix := englishHelloPrefix

	if language == spanish {
		prefix = spanishHelloPrefix
	} else if language == french {
		prefix = frenchHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}
