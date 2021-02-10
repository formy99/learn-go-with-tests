package main

import "math"

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
	//want := math.Sqrt(10.00) * math.Pi
	//print(math.(10.00) * math.Pi)
	//math.Abs(10.00)
	print(math.Abs(10.00))
}
