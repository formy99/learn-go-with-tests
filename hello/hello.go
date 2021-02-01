package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bejour, "

func Hello(name, languege string) string {

	prefix := languagueSwitcher(languege)

	if name == "" {
		name = "world"
	}
	return prefix + name
}

func languagueSwitcher(language string) (prefix string) {

	switch language {
	case "English":
		prefix = englishHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	default:
		prefix = "Hello, "
	}
	return
}
func main() {
	fmt.Println(Hello("Chriss", "English"))

}
