package main

const (
	englishHello = "Hello "
	spanishHello = "Hola "
	frenchHello  = "Bonjour "

	smile = " :)"

	spanish = "Spanish"
	french  = "French"

	defaultName = "World"
)

func Hello(name, lang string) string {
	if name == "" {
		name = defaultName
	}
	return greeting(lang) + name + smile
}

func greeting(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHello
	case spanish:
		prefix = spanishHello
	default:
		prefix = englishHello
	}
	return
}

func main() {
	//fmt.Println(Hello("world"))
}
