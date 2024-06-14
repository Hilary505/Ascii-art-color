package main

import (
	"flag"
	"fmt"
	"os"

	a "ascii/ascii_art"
)

func main() {
	// Default arguments
	input := ""
	bannerFile := "standard.txt"
	subString := " "

	flgColor := flag.String("color", "reset", "Color")
	flag.Parse()
	color := a.ColorPicker(*flgColor)
	// println("Color: ", *flgColor)

	args := flag.Args()
	nArgs := len(args)

	switch nArgs {
	case 1:
		input = args[0]
	case 2:
		if flag.NFlag() == 0 {
			input = args[0]
			bannerFile = args[1]
		} else {
			subString = args[0]
			input = args[1]
		}
	case 3:
		if flag.NFlag() == 1 {
			subString = args[0]
			input = args[1]
			bannerFile = args[2]
		}
	default:
		fmt.Println(`Usage: go run . --color=<color> <substring to be colored> "string to be printed" <banner file>`)
		os.Exit(1)
	}

	contents, err := a.GetFile(bannerFile)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Print(a.ProcessInput(contents, input, color, subString))
}
