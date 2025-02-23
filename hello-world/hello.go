package main

import "fmt"

const frenchLanguage = "fr";
const spanishLanguage = "es";


const englishHelloPrefix = "Hello, ";
const spanishHelloPrefix = "Hola, ";
const frenchHelloPrefix = "Bonjour, ";

func Hello(name, language string) string {

	if name == "" {
		name = "World";
	}

	if language == spanishLanguage {
		return spanishHelloPrefix + name + "!";
	}

	if language == frenchLanguage {
		return frenchHelloPrefix + name + "!";
	}

	return englishHelloPrefix + name + "!";
}

func main() {
	fmt.Println(Hello("Chris", "en"));
}