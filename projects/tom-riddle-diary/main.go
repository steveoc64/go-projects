package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func writeToDiary(output string) {
	file, err := os.Create("diary.txt")
	if err != nil {
		log.Fatal("Could not create file")
	}

	_, err = io.WriteString(file, output)
	if err != nil {
		log.Fatal("Failed to write to file")
	}

}

func readFromDiary(input string) string {
	data, err := ioutil.ReadFile("diary.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}

	return string(data)
}

func main() {
	fmt.Println("The book is opening...")

	const passes = 4
	OutputText := [passes]string{
		"Hello and welcome to my diary!\nDo you want to hear a story?",
		"Wonderful!\nOnce up on a time I was a student at Hogwarts, and I wrote down my thoughts in this diary.\nNow I pass it along to you. Are you ready?",
		"Well then, let us begin!\nMy name is Tom Riddle and you have found my diary.\nYou will help me get revenge! What do you think of that?",
		"I'm afraid that you don't have a choice, my dear.\nYou will help me kill Harry Potter!",
	}
	InputText := [passes]string{
		"Yes",
		"Yes, of course I am",
		"I don't wanna help you!",
		"I will obey",
	}

	for i := 0; i < passes; i++ {
		fmt.Println("Pass number:", i+1)

		fmt.Println("Writing to diary...")
		writeToDiary(OutputText[i])
		time.Sleep(5 * time.Second)

		writeToDiary("")
		fmt.Println("Waiting for input...")

		time.Sleep(15 * time.Second)
		data := readFromDiary(InputText[i])

		if InputText[i] == data {
			continue
		} else {
			log.Fatal("Invalid answer")
		}

	}

	fmt.Println("The book is closing...")
}
