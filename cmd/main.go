package main

import (
	"flag"
	"fmt"
	"thedekk/AIReview/internal/handlers"
)

func main () {
	fmt.Println("Hello, World!")

	CurrentBranch := flag.String("cb", "", "Current Branch")
	MainBranch := flag.String("mb", "main", "Main Branch")
	CustomPromt := flag.Bool("cp", false, "Custom Promt")
	OutFile := flag.String("o", "a.txt", "Out File")
	var SupplementationPromtString string
	flag.StringVar(&SupplementationPromtString,"sp", "", "Supplementation Promt")

	flag.Parse()

	fmt.Println(*CurrentBranch, *MainBranch, *CustomPromt, *OutFile, SupplementationPromtString)

	if err :=	handlers.Request(); err != nil {
		fmt.Println("Error:", err)
	}
}
