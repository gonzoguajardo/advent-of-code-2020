package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	m := make(map[int]int)
	for scanner.Scan() {
		currentString := scanner.Text()
		currentValue , _ := strconv.Atoi(currentString)
		expectedValue := 2020 - currentValue
		_, ok := m[currentValue]
		if ok {
			log.Println(currentValue * expectedValue)
			break
		}
		m[expectedValue] = currentValue
	}
}
