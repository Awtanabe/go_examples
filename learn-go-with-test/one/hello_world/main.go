package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour,"


func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	// if language == "Spanish" {
	// 	return spanishHelloPrefix + name
	// }

	// if language == "French" {
	// 	return frenchHelloPrefix + name
	// }
	// return englishHelloPrefix + name
	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}