package main

import "fmt"

// const englishHelloPrefix = "Hello "
// const spanishHelloPrefix = "Hola "
// const EnglishLang = "English"
// const SpanishLang = "Spanish"

// func Hello(name string, language string) string {
// 	prefix := ""

// 	if language == EnglishLang {
// 		prefix = englishHelloPrefix
// 		if name == "" {
// 			return prefix + "World"
// 		}
// 	} else if language == SpanishLang {
// 		prefix = spanishHelloPrefix
// 		if name == "" {
// 			return prefix + "Elodie"
// 		}
// 	}
// 	return prefix + name
// }

const (
	englishHelloPrefix = "Hello "
	spanishHelloPrefix = "Hola "
	defaultEnlishName = "World"
	defaultSpanishName = "Elodie"
)

func Hello(name, language string) string {

	if name == "" {
		switch language {
		case "Spanish":
			name = defaultSpanishName
		default:
			name =defaultEnlishName
		}
	}

	prefix := map[string]string {
		"Spanish": spanishHelloPrefix,
		"English": englishHelloPrefix,
	}[language]

	if prefix == "" {
		prefix = englishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}