package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const English = "Hello, "
const Spanish = "Hola, "
const French = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	switch language {
	case spanish:
		return Spanish
	case french:
		return French
	default:
		return English
	}
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
