package main

import "fmt"

const (
	frenchLanguage = "fr";
	spanishLanguage = "es";

	englishHelloPrefix = "Hello, ";
	spanishHelloPrefix = "Hola, ";
	frenchHelloPrefix = "Bonjour, ";
)

func Hello(name, language string) string {

	if name == "" {
		name = "World";
	}

	return greetingPrefix(language) + name + "!";
}

// different ways to handle return in Go
// 1. prefix - named return value
// 2. return;
func greetingPrefix(language string) (prefix string) {
	switch language {
	case frenchLanguage:
		return frenchHelloPrefix;
	case spanishLanguage:
		return spanishHelloPrefix;
	default:
		prefix = englishHelloPrefix;
	}

	return;
}

// public functions start with CapitalLetter and private functions with lowercaseLetter

func main() {
	fmt.Println(Hello("Chris", "en"));
}