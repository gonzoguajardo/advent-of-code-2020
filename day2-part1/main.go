package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	totalPassed := 0
	for scanner.Scan() {
		currentString := scanner.Text()
		spaceSplit := strings.Split(currentString, " ")
		hyphenSplit := strings.Split(spaceSplit[0], "-")
		min , _:= strconv.Atoi(hyphenSplit[0])
		max , _:= strconv.Atoi(hyphenSplit[1])
		letter := spaceSplit[1][0:1]
		checkString := spaceSplit[2]
		count := strings.Count(checkString, letter)
		pass := count >= min && count <= max
		if pass {
			totalPassed++
		}
	}
	log.Println(totalPassed)
}
