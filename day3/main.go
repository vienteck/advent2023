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
	result := 0
	startOfNumber := ""
	addNumber := false
	for j := 0; j < len(lineText); j++ {
		for i := 0; i < len(lineText[j]); i++ {
			if string(lineText[j][i]) >= "0" && string(lineText[j][i]) <= "9" {
				if startOfNumber == "" {
					addNumber = false
					//check above
					if j > 0 {
						if string(lineText[j-1][i]) != "." {
							addNumber = true
						}
					}
					//check left
					if i > 0 {
						if string(lineText[j][i-1]) != "." {
							addNumber = true
						}
					}
					//check upper-left
					if j > 0 && i > 0 {
						if string(lineText[j-1][i-1]) != "." {
							addNumber = true
						}
					}
					//check down
					if j < len(lineText)-1 {
						if string(lineText[j+1][i]) != "." {
							addNumber = true
						}
					}
					//check lower-left
					if j < len(lineText)-1 && i > 0 {
						if string(lineText[j+1][i-1]) != "." {
							addNumber = true
						}
					}
					//check left, up-left and down-left
				} else {
					//check down
					if j < len(lineText)-1 {
						if string(lineText[j+1][i]) != "." {
							addNumber = true
						}
					}
					//check above
					if j > 0 {
						if string(lineText[j-1][i]) != "." {
							addNumber = true
						}
					}
				}

				startOfNumber = fmt.Sprintf("%v%v", startOfNumber, string(lineText[j][i]))
			} else {
				if startOfNumber != "" {
					number, _ := strconv.Atoi(startOfNumber)
					//check down
					if j < len(lineText)-1 {
						if string(lineText[j+1][i]) != "." {
							addNumber = true
						}
					}
					//check above
					if j > 0 {
						if string(lineText[j-1][i]) != "." {
							addNumber = true
						}
					}
					if string(lineText[j][i]) != "." {
						addNumber = true
					}
					if addNumber {
						result += number
					} else {
						fmt.Printf("Not printing nubmer %v line %v\n", number, j)
					}
					startOfNumber = ""
				}
			}

			if i == len(lineText[j])-1 {
				if addNumber == true {
					number, _ := strconv.Atoi(startOfNumber)
					result += number
				}
			}
		}
		fmt.Println(result)
	}
}
