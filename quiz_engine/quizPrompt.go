package quizengine

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func handleError(err error){
	if err != nil {
		log.Panic(err)
	}
}


func getFiles(dirPath string) []string {
	var fileNames []string;
	files, err := ioutil.ReadDir(dirPath)
	handleError(err)
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames
}


func InitalPrompt() {

	var userInputNumber int
	quizDirName := "/quiz_problems"
	fmt.Println("Welcome! \nWhich quiz would you like to take?")
	wd, err := os.Getwd()
	handleError(err)
	var fileNames = getFiles(wd + quizDirName)
	for i, file := range fileNames {
		fmt.Println(strconv.Itoa(i+1) + ": " + file[:len(file) -4] + " ?")
	}
	for {
		fmt.Print("Enter quiz number: ")
		fmt.Scanf("%d", &userInputNumber)
		if userInputNumber -1 < len(fileNames){
			runQuiz(fileNames[userInputNumber - 1], wd + quizDirName)
			break;
		}else{
			fmt.Println("Number doesn't exist, try again")
		}
	}
}

func countCSV(file *csv.Reader) int{
	temp := 0
	for {
		_, err := file.Read()
			if err == io.EOF {
				break
			}
			temp++
	}
	return temp
}


func runQuiz(quiz_name string, workDir string) {
	var score int
	var currInputAns string
	problemNumber := 0
	filePath := workDir +"/"+ quiz_name
	file, err := os.Open(filePath)
	handleError(err)
	csvFile := csv.NewReader(file);
	timeoutChannel := make(chan bool)
	go func(){
		for {
		record, err := csvFile.Read()
			if err == io.EOF {
				timeoutChannel <- true
			}
		question := record[0]
		answer := record[1]
		fmt.Print("Problem #"+ strconv.Itoa(problemNumber) + " " + question + " : ")
			fmt.Scanf("%s", &currInputAns)
			currInputAns = strings.TrimSpace(currInputAns)
			if currInputAns == answer {
				score++
			}
			problemNumber++
		}
	}()
	select {
	case <- time.After(time.Second * 30):
		fmt.Println("TIMEOUT")
		problemNumber += countCSV(csvFile)
	case <- timeoutChannel:
		fmt.Println("You scored " + strconv.Itoa(score) + " out of " + strconv.Itoa(problemNumber))
		os.Exit(0);
		
	}
	fmt.Println("You scored " + strconv.Itoa(score) + " out of " + strconv.Itoa(problemNumber))
}

