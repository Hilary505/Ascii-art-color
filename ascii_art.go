package ascii_art

import (
	"fmt"
	"os"
	"strings"
)

// GetFile reads from the file specified by filename and returns its contents
func GetFile(filename string) ([]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("An error", err)
		os.Exit(0)
	}

	if len(file) == 0 {
		fmt.Println("Error: The bunner file is empty")
		os.Exit(1)
	}

	myfile := string(file)
	var contents []string

	if filename == "thinkertoy.txt" {
		contents = strings.Split(myfile, "\r\n")
	} else {
		contents = strings.Split(myfile, "\n")
	}

	return contents, nil
}

func ColorPicker(color string) (colorCode string) {
	colorChat := map[string]string{
		"reset":   "\u001b[39m",
		"red":     "\u001b[31m",
		"green":   "\u001b[32m",
		"magenta": "\u001b[35m",
	}

	 _, ok := colorChat[color] 

	 if !ok {
		fmt.Println("Error: Color Not Found")
		os.Exit(1)
	 }
	
	colorCode = colorChat[color]

	return colorCode
}

// ProcessInput accepts the contents of the ASCII art file and the input string,
// and processes the input to display the corresponding ASCII art
func ProcessInput(contents []string, input, color, subString string) (strArt string) {
	count := 0
	input = strings.ReplaceAll(input, "\n", "\\n")
	input = strings.ReplaceAll(input, "\\t", "    ")
	newInput := strings.Split(input, "\\n")

	ci := -1

	for _, arg := range newInput {
		if arg == "" {
			count++
			if count < len(newInput) {

				strArt += "\n"

				continue
			} else {
				continue
			}
		}

		for i := 1; i <= 8; i++ {
			if subString != "" && subString != " " {
				ci = strings.Index(arg, subString)
			}

			if subString ==" " {
				ci = 0
			}

			for j, ch := range arg {
				if ch > 126 {
					fmt.Println("The strihs contains an unprintable character", ch)
					os.Exit(0)
				}

				index := int(ch-32)*9 + i

				if index >= 0 && index < len(contents) {
					if ci == j {
						strArt += color
					}

					strArt += (contents[index])

					if ci != -1 && ci+len(subString)-1 == j && j < len(arg)-1 && subString != "" && subString != " " {
						strArt += ColorPicker("reset")
						ci = strings.Index(arg[j+1:], subString) + j + 1
					}

				}
			}
			strArt += ColorPicker("reset") + "\n"
		}
	}

	return 
}
