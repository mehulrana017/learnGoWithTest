package main

import "fmt"

var prefix = "hello, "

func hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	if lang == "spanish" {
		return "hola, " + name
	}
	return prefix + name
}

func main() {
	fmt.Println(hello("Chris", ""))
}
