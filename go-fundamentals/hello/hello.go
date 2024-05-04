package hello

import "fmt"

const (
	spanish = "ES"
	french  = "FR"
	german  = "DE"

	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Hola"
	frenchHelloPrefix  = "Bonjour"
	germanHelloPrefix  = "Guten tag"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + ", " + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix

	case french:
		prefix = frenchHelloPrefix

	case german:
		prefix = germanHelloPrefix

	default:
		prefix = englishHelloPrefix

	}

	return
}

func main() {
	fmt.Println(Hello("Bob", ""))
}
