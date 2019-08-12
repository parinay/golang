package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Quiz struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func main() {
	csvFile, err := os.Open("quiz.csv")
	if err != nil {
		log.Fatal("error:", err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var quiz []Quiz

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("error:", err)
		}

		quiz = append(quiz, Quiz{
			Question: record[0],
			Answer:   record[1],
		})
	}
	quizJeson, err := json.MarshalIndent(quiz, "", "	")
	if err != nil {
		log.Fatal("error:", err)
	}
	fmt.Println(string(quizJeson))
}
