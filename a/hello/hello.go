package hello

import "rsc.io/quote"

// GreetWorld ...
func GreetWorld() string {
	return quote.Hello()
}

// GreetUser ..,,.
func GreetUser(name string) string {
	if name == "" {
		return GreetWorld()
	}
	return "Hello, " + name + "."
}
