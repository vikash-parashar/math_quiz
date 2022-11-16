package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type Quiz struct {
	question string
	answer   string
}

func main() {
	csvFileName := flag.String("csv", "quiz.csv", "a csv file which contain questions and answers")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		log.Println("failed to open that csv file")
		os.Exit(1)
	}

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln("failed to read the data from csv file")
	}
	fmt.Println(records)

	for i, v := range records {
		fmt.Println(i)
		fmt.Println(v)
	}

}
