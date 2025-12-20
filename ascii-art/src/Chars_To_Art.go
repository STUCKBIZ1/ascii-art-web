package ascii_art

import (
	"fmt"
	"strings"
)

func IsLine(s []string) bool {
	for _, v := range s {
		if v != "" {
			return false
		}
	}
	return true
}

func Chars_To_Art(fonts [][]string, input string) string {
	var CInput [][]string
	var result string
	input = strings.Replace(input, "\r", "", 1)
	sliceInput := strings.Split(input, "\n")
	if IsLine(sliceInput) {
		sliceInput = sliceInput[1:]
	}

	for _, inputline := range sliceInput {
		if inputline == "" {
			fmt.Println()
			continue
		}
		for _, char := range inputline {
			for j, font := range fonts {
				if j+32 == int(char) {
					CInput = append(CInput, font)
				}
			}
		}
		result += Print_Fonts(CInput)
		CInput = [][]string{}
	}
	fmt.Println()
	return result
}
