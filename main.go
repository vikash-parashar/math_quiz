package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
)

type problem struct {
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
	// fmt.Println(records)

	// for i, v := range records {
	// 	fmt.Println(i)
	// 	fmt.Println(v)
	// }

	parseLines(records)

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
