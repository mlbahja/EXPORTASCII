package ascii_web

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArtGenerator(data, banner string) (string, int) {
	asciiFile, err := os.ReadFile("utils/" + banner + ".txt")
	if err != nil {
		fmt.Println(err)
		return "", 1
	}
	asciiCharacters := strings.ReplaceAll(string(asciiFile), "\r", "")
	asciiTable := AsciiTableMaker(asciiCharacters)
	arg := strings.ReplaceAll(data, "\r", "")

	arr := strings.Split(arg, "\n")
	for _, str := range arr {
		for _, c := range str {
			if c < ' ' || c > '~' {
				fmt.Println("Error: Character out of range.")
				return "", 1
			}
		}
	}

	if len(arg) == strings.Count(arg, "\n") {
		arr = arr[:len(arr)-1]
	}
	output := ""
	for _, str := range arr {
		if len(str) == 0 {
			output += "\n"
		} else {
			for i := 0; i < 8; i++ {
				for _, c := range str {
					if c >= ' ' && c <= '~' {
						output += asciiTable[int(c)-int(' ')][i]
					}
				}
				output += "\n"
			}
		}
	}
	return output, 0
}
