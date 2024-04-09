package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "the required language")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

// language represents the language's code
type language string

// phrasebook holds greetings for each support language
var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
	"fr": "Bonjour le monde",  // French
	"en": "Hello World",       // English
	"he": "שלום עולם",         // Hebrew
	"ur": "ہیلو دنیا",         // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
}

// greet says hello in various languages
func greet(l language) string {
	greeting, ok := phrasebook[l]

	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return greeting

}
