package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type question struct {
	question string
	answer   string
}

func main() {
	fmt.Println("Hello World")
	// Read flags
	var sFile = flag.String("f", "problems.csv", "The questions file to use.  Default is problems.csv")
	var sTimer = flag.String("t", "4", "The time limit for each question, in seconds.  Default is 10 seconds")
	flag.Parse()

	fmt.Println("File: " + *sFile)

	// Read CSV
	f, err := os.Open(*sFile)

	if err != nil {
		log.Fatal(err)
	}

	var sQuestions [](question)

	r := csv.NewReader(f)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		//how do I check the record length and type?

		var sQuestion question
		sQuestion.question = record[0]
		sQuestion.answer = record[1]

		sQuestions = append(sQuestions, sQuestion)
	}

	var nNumQuestions, nCorrect int

	nNumQuestions = len(sQuestions)

	fmt.Printf("Push enter to begin!")
	fmt.Scanln()

	go func() {
		var nTimer, _ = strconv.Atoi(*sTimer)
		time.Sleep(time.Duration(nTimer) * time.Second)
		fmt.Println("you're out of time!")
		fmt.Printf("You got %v of %v correct!\n", nCorrect, nNumQuestions)
		os.Exit(0)
	}()

	//Ask questions
	for _, v := range sQuestions {
		fmt.Println("What is " + v.question + "?")
		var answer string

		fmt.Scanln(&answer)

		if answer == v.answer {
			fmt.Println("correct!")
			nCorrect++
		} else {
			fmt.Println("incorrect!")
		}
	}

	//Display score
	fmt.Printf("You got %v of %v correct!\n", nCorrect, nNumQuestions)

}
