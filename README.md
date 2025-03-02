# clip

`clip` is a small, opinionated argument parser for Golang.

## Usage

```go
package main

import "fmt"
import "os"

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
        os.Exit(0)
	}
    
	if c.NArgs() != 1 {
		fmt.Println("Error: Must provide the name to greet")
		os.Exit(2)
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
```

Then, invoking the application without any arguments yields:

```
Error: Must provide the name to greet

greeter: Greats someone
usage: greeter [--version | -v] [--language value | -l value] [--help | -h]  NAME
               

Copyright (c) Vito Sartori
Author: Vito Sartori <hey@vito.io>
Version: v0.1.0 (c0ffeebabe)

Flags:
    [--version | -v] Print the version information 
    [--language value | -l value] Sets the language 
    [--help | -h] Shows this message and exits 


(Process finished with the exit code 2)
```

With `-v`:

```
greeter version v0.1.0

(Process finished with the exit code 0)
```

With `-l jp Vito`

```
こんにちは Vito!

(Process finished with the exit code 0)
```

## License

```
The MIT License (MIT)

Copyright (c) Vito Sartori

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```
