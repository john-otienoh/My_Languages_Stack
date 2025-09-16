package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func flagUsage() {
	usageText := `example05 is an example cli tool.
	
	Usage:
	main command [arguments]
	The commands are:
		uppercase uppercase a string
		lowercase lowercase a string
	Use "main [command] --help" for more information about a command.`

	fmt.Fprintf(os.Stderr, "%s\n\n", usageText)
}
func main() {
	// Accessing Command-Line Arguments
	for i, arg := range os.Args {
		fmt.Println("Argument", i, "is", arg)
	}
	// Parsing command line flags
	// Flag Parsing for a String, Int, and Boolean
	s := flag.String("s", "Hello World", "String help text")
	i := flag.Int("i", 1, "Int help text")
	b := flag.Bool("b", false, "Bool help text")

	flag.Parse()

	fmt.Println("Value of s:", *s)
	fmt.Println("value of i:", *i)
	fmt.Println("value of b:", *b)

	// Creating Help Text for a Command-Line Tool
	flag.Usage = func() {
		usageText := `Usage main [OPTION] An example of customizing usage output
		-s, --s	example string argument, default: String help text
		-i, --i	example integer argument, default: Int help text
		-b, --b	example boolean argument, default: Bool help text`
		fmt.Fprintf(os.Stderr, "%s\n", usageText)
	}

	// Creating Subcommands
	flag.Usage = flagUsage
	uppercaseCmd := flag.NewFlagSet("uppercase", flag.ExitOnError)
	lowercaseCmd := flag.NewFlagSet("lowercase", flag.ExitOnError)

	if len(os.Args) == 1 {
		flag.Usage()
		return
	}
	switch os.Args[1] {
	case "uppercase":
		s := uppercaseCmd.String("s", "", "A string of text to be uppercased")
		uppercaseCmd.Parse(os.Args[2:])
		fmt.Println(strings.ToUpper(*s))
	case "lowercase":
		s := lowercaseCmd.String("s", "", "A string of text to be lowercased")
		lowercaseCmd.Parse(os.Args[2:])
		fmt.Println(strings.ToLower(*s))
	}
}
