package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"ascii/ascii_art"
)

func main() {
	// Default arguments
	input := ""
	bannerFile := "shadow.txt"
	subString := " "

	var flgColor = flag.String("color", "reset", "Color")
	var flagOutput = flag.String("output", "", "my bannerFiele")
	flag.Parse()
	color := ascii_art.ColorPicker(*flgColor)
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
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println()

		fmt.Println("EX: go run . something standard")
		os.Exit(0)
	}
	if !strings.HasSuffix(bannerFile, ".txt") {
		bannerFile += ".txt"
	}

	contents, err := ascii_art.GetFile(bannerFile)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	output := ascii_art.ProcessInput(contents, input, color, subString)

	if *flagOutput != "" {
		err = os.WriteFile(*flagOutput, []byte(output), 0644)
		if err != nil {
			fmt.Println("An erreor just occured while writing File:")
			os.Exit(1)
		}
	}else{
		fmt.Print(output)
	}

}

// var fame = flag.String("output", "banner.txt", "my file input")
// 	flag.Parse()
// 	args1 := flag.Args()
// 	text := args1[0] + ("\n")
// 	err = os.WriteFile(*fame, []byte(text), 0644)
// 	if err != nil {
// 		return
// 	}
