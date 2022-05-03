// https://github.com/gophercises/quiz
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	const TimeLimit time.Duration = 10

	reader := bufio.NewReader(os.Stdin)

	csvfile, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
		os.Exit(1)
	}

	r := csv.NewReader(csvfile)

	correctCount, wrongCount := 0, 0

	/************************************************/
	var dummy string
	fmt.Println("Press a key to start quiz !")
	_, _ = fmt.Scanln(&dummy)
	/************************************************/

	timer := time.NewTimer(time.Second * TimeLimit)

	func() {
		for {
			select {
			case <-timer.C:
				fmt.Println("test is over!")
				return
			default:
				record, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Question is: %s, your answer ? ", record[0])
				answer, _ := reader.ReadString('\n')

				if answer = strings.Replace(answer, "\n", "", -1); strings.Compare(answer, record[1]) == 0 {
					fmt.Println("Correct Answer! Next question:")
					correctCount++
				} else {
					fmt.Println("Wrong Answer! Next question:")
					wrongCount++
				}
			}
		}
	}()

	fmt.Printf("The test is done. Correct answers: %v, Wrong Answers: %v\n", correctCount, wrongCount)
}
