package main

import "fmt"

import "github.com/heyvito/clip"

func main() {
	app := clip.New(
		clip.Name("greeter"),
		clip.Description("Greats someone"),
		clip.Copyright("(c) Vito Sartori"),
		clip.Author("Vito Sartori <hey@vito.io>"),
		clip.Version("v0.1.0"),
		clip.Build("c0ffeebabe"),
		clip.ArgsDescription("NAME"),
		clip.Entrypoint(entrypoint),
	)
	app.Boolean("version").Shorthand("v").Description("Print the version information")
	app.Options("language", "en", "it", "es", "pt", "jp").
		Shorthand("l").
		Description("Sets the language").
		Default("en")
	app.TakeArguments()
	app.Run()
}

func entrypoint(c *clip.CLI) {
	if c.Boolean("version") {
		fmt.Println("greeter version v0.1.0")
		return
	}

	if c.NArgs() != 1 {
		fmt.Print("Error: Must provide the name to greet\n\n")
		c.PrintHelpExit(2)
	}

	switch c.Option("language") {
	case "en":
		fmt.Print("Hello, ")
	case "it":
		fmt.Print("Ciao, ")
	case "es":
		fmt.Print("¡Hola, ")
	case "pt":
		fmt.Print("Olá, ")
	case "jp":
		fmt.Print("こんにちは ")
	}

	fmt.Printf("%s!\n", c.Arg(0))
}
