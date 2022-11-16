package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type problem struct {
	question string
	answer   string
}

var count int

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
	// fmt.Println(records)

	// for i, v := range records {
	// 	fmt.Println(i)
	// 	fmt.Println(v)
	// }

	problems := parseLines(records)
	var ans string
	for i, v := range problems {
		fmt.Printf("Problem No : %d \n\t  %s : \n", i+1, v.question)
		fmt.Scanf("%s\n", &ans)
		if v.answer == ans {
			fmt.Println("---->  Right Answer  <----")
			count++
		} else {
			fmt.Println("---->  Wrong Answer  <----")
		}
	}

	fmt.Printf("Your Score Is %d Out Of %d ", count, len(problems))

}

func parseLines(lines [][]string) []problem {

	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return ret

}
