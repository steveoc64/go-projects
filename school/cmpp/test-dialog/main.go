package main

import (
	"fmt"

	"github.com/sqweek/dialog"
)

func main() {
	filename, err := dialog.File().Filter("PDF file", "pdf").Load()
	if err != nil {
		panic(err)
	}

	fmt.Println(filename)
}
