package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineText := []string{}
	for scanner.Scan() {
		row := scanner.Text()
		lineText = append(lineText, row)
	}

	startOfNumber := ""
	for j := 0; j < len(lineText); j++ {
		for i := 0; i < len(lineText[j]); i++ {
			addNumber := true
			if string(lineText[j][i]) > "0" && string(lineText[j][i]) < "9" {
				if startOfNumber == "" {
					//check above
					if j > 0 {
						if string(lineText[j-1][i]) == "." {
							addNumber = false
						}
					}
					//check left
					if i > 0 {
						addNumber = false
						fmt.Println(addNumber)
					}
					//check left, up-left and down-left
				} else {
					//check up and down
				}

				startOfNumber = fmt.Sprintf("%v%v", startOfNumber, string(lineText[j][i]))
			} else {
				if startOfNumber != "" {
					//process number
					number, _ := strconv.Atoi(startOfNumber)
					fmt.Println(number)
					startOfNumber = ""
				}
			}
		}
	}

}
